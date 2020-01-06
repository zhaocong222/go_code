package main

import (
	"fmt"
)

func main() {

	//可以从控制台接收用户信息，【姓名、年龄、薪水】
	//方式1 fmt.Scanln
	//1.先声明需要的变量
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)

	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Println("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)



}