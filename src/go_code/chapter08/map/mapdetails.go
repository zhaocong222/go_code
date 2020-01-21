package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 1232
}

func main() {
	//map是引用类型，遵守引用类型传递的机制，在一个函数接受map
	//修改后，会直接修改原来的map
	//map可以自动扩容，这里给的最大空间为2
	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2

	fmt.Println(map1)
	modify(map1)
	fmt.Println(map1)

}
