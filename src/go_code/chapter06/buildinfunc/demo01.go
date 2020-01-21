package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("num1类型是%T,num1的值是%v,num1的地址是%v\n", num1, num1, &num1)

	num2 := new(int) // *int 指针类型
	//num2的类型 => *int
	//num2的值  => 指向某个空间地址的值
	//num2的地址  => 本身的地址
	//num2地址指向的值
	*num2 = 112
	fmt.Printf("num2类型是%T,num2的值是%v,num2的地址是%v, num2指向的值是%v\n", num2, num2, &num2, *num2)

}
