package main

import (
	"fmt"
)

func main()  {
	//赋值运算符的使用演示
	//var i int
	//i = 10 

	//有两个变量，a和b，要求将其进行交换，最终打印结果
	//a = 9, b = 2 ==> a = 2 b = 9
	a := 9
	b := 2
	//定义一个临时变量,交换
	t := a
	a = b
	b = t
	fmt.Printf("交换后的情况是 a = %v, b = %v\n", a, b)

	//复合赋值的操作
	a += 17 //等价 a = a + 17
	fmt.Println("a=", a)

	var c int
	c = a + 3
	fmt.Println("c=", c)





}