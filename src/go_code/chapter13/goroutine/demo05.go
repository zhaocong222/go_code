package main

import "fmt"

func main() {

	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	intChan <- 300
	//关闭了管道后，不能再写入，但是可以继续读取
	close(intChan)

	//测试写入 panic: send on closed channel
	//intChan <- 500

	//循环读取,如果在channle没有关闭，则会出现deadlock的错误
	for v := range intChan {
		fmt.Println("v=", v)
	}

}
