package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (s Stu) Say() {
	fmt.Println("学生say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i = ", i)
}

func main() {
	var stu Stu //结构体变量，实现了Say()，实现了Interface
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say()
}
