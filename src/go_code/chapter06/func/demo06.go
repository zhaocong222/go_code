package main

import "fmt"

var (
	//fun1就是一个全局匿名函数
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 求2个数的和，使用匿名函数
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1=", res1)

	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	//a的类型是函数类型
	fmt.Printf("a的类型是%T,调用a的返回值是%v\n", a, a(2, 1))

	//调用全局匿名函数
	res4 := fun1(4, 9)
	fmt.Println("res4=", res4)

}
