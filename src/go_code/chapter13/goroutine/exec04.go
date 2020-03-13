package main

import (
	"fmt"
	"time"
)

//向intChan放入 1-8000个数
func putNum(intChan chan int) {

	for i := 1; i <= 10; i++ {
		intChan <- i
	}
	close(intChan)
}

//从管道中读取
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//使用for循环
	var flag bool
	for {
		time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok {
			//取不到东西
			break
		}

		//判断素数
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//说明不是素数
				flag = false
				break
			}
		}

		if flag {
			//是素数,放入primeChan
			primeChan <- num
		}
	}

	fmt.Println("有一个primeChan协程因为取不到数据，退出")
	//这里我们还不能关闭primeChan
	//向 exitChan 写入true
	exitChan <- true

}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) //放入结果
	//标识退出的管道
	exitChan := make(chan bool, 4) //4个标识

	//开启一个协程，向intChan放入 1-8000个数
	go putNum(intChan)
	//开启4个协程，从intChan取出数据，并判断是否为素数，如果是就放入到primeChan
	for i := 0; i < 8000; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//协程匿名函数
	go func() {

		for i := 0; i < 4; i++ {
			<-exitChan //阻塞
		}

		//当我们从exitChan取出了4个结果，就可以放心的关闭primeChan
		close(primeChan)
		close(exitChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}

		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main主线程退出")
}
