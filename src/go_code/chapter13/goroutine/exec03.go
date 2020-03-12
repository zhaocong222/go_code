package main

import (
	"fmt"
	"time"
)

//write data
func writeData2(intChan chan int) {
	//管道只有20容量，但是要写入50个元素,编译器检测到没有读取的协程，会产生死锁
	//如果写的很快，但是读的很慢，会发送死锁吗？
	//不会，编译器只要检测到管道有协程的读取，就不会发生死锁，只是会控制写入变慢 (编译器底层的管道阻塞机制)
	for i := 0; i < 50; i++ {
		fmt.Printf("writeData 写入数据=%v\n", i)
		intChan <- i
	}
	//写完后关闭管道，不影响读
	close(intChan)
}

//read data
func readData2(intChan chan int, exitChan chan bool) {
	for {
		//每休眠1秒读1次
		time.Sleep(time.Second)
		if v, ok := <-intChan; ok {
			fmt.Printf("readData 读到数据=%v\n", v)
		} else {
			break
		}
	}
	//读完了,任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	//创建两个管道
	intChan := make(chan int, 20)
	exitChan := make(chan bool, 1)

	go writeData2(intChan)

	go readData2(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

	fmt.Println("协程完成工作，主线程执行完毕")

}
