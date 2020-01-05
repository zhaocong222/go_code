package main 

import (
	"fmt"
)

//指针类型
func main() {

	//基本数据类型
	i := 10
	//i的地址
	fmt.Println("i的地址=", &i)

	//含义
	//1. ptr是一个指针变量
	//2. ptr的类型是 *int
	//3. ptr 本身的值是&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	//获取指针类型的值
	fmt.Printf("ptr指向的值是 %v\n", *ptr)

}