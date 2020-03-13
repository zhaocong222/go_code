package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func test() {
	//这里我们可以使用defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang" //error,这里还没make

}

func main() {
	//2个协程，但是其中一个出现了panic,如果我们没有捕获这个panic
	//就会造成整个程序崩溃，这时我们可以在协程种使用defer + recover捕获panic处理
	//这样出问题的协程不会影响主线程和其他协程
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}

}
