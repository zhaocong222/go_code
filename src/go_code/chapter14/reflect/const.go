package main

import "fmt"

func main() {
	var num int
	//常量声明时必须给值,只能修饰bool，数值类型(int,float系列)，string类型
	const tax int = 0
	const b = 9 / 3
	fmt.Println(num, tax, b)

	const (
		c = iota //表示初始化为0
		d        //c的基础上+1
		e        //d的基础上+1
	)
	fmt.Println(c, d, e)
}
