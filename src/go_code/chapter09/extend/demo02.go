package main

import "fmt"

type A struct {
	a1   string
	A2   string
	Name string
}

type B struct {
	b1   string
	B2   string
	Name string
	A
}

func main() {
	//在golang中，不管开头大写/小写，在包内(该文件内)都可以访问
	var b B
	b.b1 = "b1"
	b.B2 = "B2"
	fmt.Println(b.b1, b.B2)

	b.A.a1 = "a1"
	b.A.A2 = "A2"

	fmt.Println(b.A.a1, b.A.A2)

	//也可以直接简易写法,
	//编译器会先查找b结构体如果没有a1属性，再查找嵌入结构体A里面是否有此属性
	b.a1 = "a3"
	b.A2 = "A4"
	b.Name = "kuku"
	b.A.Name = "koo"

	fmt.Println(b.a1, b.A2, b.Name, b.A.Name)

}
