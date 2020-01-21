package main

import "fmt"

func main() {
	//演示切片的使用make, 5个元素，10个容量
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	//对于切片，必须make使用
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice)) //容量

	//方式3
	var slice2 []string = []string{"tome", "jack", "mary"}
	fmt.Println(slice2)
	fmt.Println("slice2的size=", len(slice2))
	fmt.Println("slice2的cap=", cap(slice2)) //容量

}
