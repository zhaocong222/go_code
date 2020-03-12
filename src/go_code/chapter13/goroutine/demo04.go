package main

import "fmt"

func main() {
	//演示一下管道的使用,chan队列
	//1.创建一个可以存在3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2.看看intChan是什么,intChan是指向管道的地址
	fmt.Printf("intChan=%v, intChan本身的地址是%v\n", intChan, &intChan)

	//3.像管道写入数据,<-写入， ->读取
	intChan <- 10
	num := 211
	intChan <- num

	//注意：当我们给管道写入数据时，不能超过其容量,否则报错

	//4.输出管道的长度和cp(容量,一般make时候)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	//5.从管道取出数据
	num2 := <-intChan
	fmt.Println("num2=", num2)

	//6.在没有使用协程的情况下，如果我们的管道数据已经全部取完，再取就会报告defalock
	num3 := <-intChan
	num4 := <-intChan

	fmt.Println("num3", num3, "num4", num4)

}
