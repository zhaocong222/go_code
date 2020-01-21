package main

import "fmt"

//先执行
var age = test()

//为了看到全局变量先被初始化的，我们这里先写函数初始化
func test() int {
	fmt.Println("test()")
	return 90
}

//可以再init中完成初始化工作
func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()...", age)
}
