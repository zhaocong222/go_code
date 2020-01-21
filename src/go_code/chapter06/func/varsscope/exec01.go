package main

import "fmt"

var name = "tom" //全局变量
//Age := 20  //函数外不能这么写

func test01() {
	fmt.Println(name) //tom
}

func test02() {
	name := "jack"
	fmt.Println(name) //jack
}

func main() {
	fmt.Println(name) //tom
	test01()
	test02()
	test01()

}
