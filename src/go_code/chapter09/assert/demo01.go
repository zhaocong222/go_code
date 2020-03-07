package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	point := Point{1, 2}
	a = point
	//如何将a赋给一个point变量
	var b Point
	//b = a //不可以这么写
	b = a.(Point) //类型断言
	fmt.Println(b)

	//类型断言的其他案例
	var x interface{}
	var c float32 = 1.1
	x = c //空接口，可以接收任何类型
	//x=>float32 [使用类型断言]
	//y := x.(float32)

	//代码崩溃，使用类型断言
	y, ok := x.(float64)
	if ok {
		fmt.Printf("y的类型是%T,值是%v", y, y)
		fmt.Println("okokokokook!")
	} else {
		fmt.Println("转换失败")
	}

}
