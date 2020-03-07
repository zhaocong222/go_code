package main

import "fmt"

type Test struct {
	Age  int
	Name string
}

func (t Test) test01() {
	t.Age = 20
	t.Name = "赵云"
}

func (t *Test) test02() {
	//t是指针变量，(*t).Age
	t.Age = 30 //底层会自动优化成(*t).Age
	t.Name = "司马懿"
}

func main() {

	te := Test{
		Age:  1,
		Name: "韩德",
	}
	fmt.Println("原始数据")
	fmt.Println(te)

	fmt.Println("调用test01")
	te.test01()
	(&te).test01()
	fmt.Println(te)

	fmt.Println("调用test02")
	te.test02() //等价(&te).test02()
	//(&te).test02() //应该是这个写，但是te.test02()写法go底层编译器会自动优化成地址传递
	fmt.Println(te)
}
