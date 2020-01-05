package main

import (
	"fmt"
	"strconv"
)

//基本数类型转成string
func main()  {
	var num1 int = 99
	var num2 float64 = 23.345
	var b bool = true
	var myChar byte = 'h'
	var str string //空的str

	//使用第一种方式转换 fmt.Sprintf方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%v\n", str, str)

	// 第二种方式 利用函数strconv函数
	var num3 int = 99
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%v\n", str, str)

}