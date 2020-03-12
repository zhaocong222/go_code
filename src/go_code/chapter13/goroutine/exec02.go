package main

import "fmt"

//write data
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		fmt.Printf("writeData 写入数据=%v\n", i)
		intChan <- i
	}
	//写完后关闭管道，不影响读
	close(intChan)
}

//read data
func readData(intChan chan int, exitChan chan bool) {
	for {
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
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

	fmt.Println("协程完成工作，主线程执行完毕")

}
