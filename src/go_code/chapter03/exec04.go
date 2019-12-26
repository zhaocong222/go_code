package main

import "fmt"


func main() {
	// 该区域的数据值可以在同一类型范围内不断变化
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i=", i)

	//变量在同一个作用域（在一个函数或者在代码块）内不能重名
	//var i int = 59  //i redeclared in this block



}