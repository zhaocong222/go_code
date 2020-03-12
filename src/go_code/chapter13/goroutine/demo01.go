package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello,world" + strconv.Itoa(i))
		//睡眠1秒
		time.Sleep(time.Second)
	}
}

func main() {

	//test() //执行完了才往下执行

	go test() //开启了协程 执行test,观察输出

	//主函数代码 , 主线程执行完退出了，协程即使还没执行完也会退出
	for i := 1; i <= 2; i++ {
		fmt.Println("hello,golang" + strconv.Itoa(i))
		//睡眠1秒
		time.Sleep(time.Second)
	}
}
