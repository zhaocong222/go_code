package main

import "fmt"

type Person struct {
	Name string
}

//给A类型绑定一份方法
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

func main() {
	var p Person
	p.Name = "tom"
	p.test() //调用方法
}
