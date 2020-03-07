package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式1
	//var p1 Person

	//方式2
	p2 := Person{"tome", 20}
	p2.Age = 20

	//方式3 var person *Person = new (Person)
	//p3是一个指针
	p3 := new(Person)
	//因为p3是一个指针，因此标准的给字段赋值方式
	//(*p3).Name = "smith"也可以这样写 p3.Name = "smith"
	//原因：go的设计者 为了程序员使用方便，底层会对 p3.Name = "smith" 进行处理
	//会给p3加上取值运算p3.Name = "smith" 在底层就是 (*p3).Name = "smith"
	(*p3).Name = "smith"
	(*p3).Age = 30
	p3.Name = "john"
	fmt.Println()

	//方式4
	var person *Person = &Person{}
	//因为person是一个指针，因此标准的访问字段的方法
	// (*person).Name = "scott"
	// go的设计者为了程序员使用方便，也可以person.Name = "scott"
	// 原因和上面一样，底层会对person.Name = "scott"进行处理，会加上(*person)
	//下面2种写法都可以
	(*person).Name = "scott"
	person.Name = "Guia"
}
