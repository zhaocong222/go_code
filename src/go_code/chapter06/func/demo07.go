package main

import (
	"fmt"
	"strings"
)

//闭包，返回的是函数类型， func(int) int
func AddUper() func(int) int {

	var n int = 10
	var str = "hello"

	return func(x int) int {
		n = n + x
		//str += string(x) //string(x) => 如果x是整数转成对应的ascii码
		str += "a"
		fmt.Println("str=", str)
		return n
	}

}

//1.编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg),并返回一个闭包
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果字符串没有指定后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func main() {
	//使用前面的代码
	f := AddUper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

	//返回一个闭包
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("winter"))

}
