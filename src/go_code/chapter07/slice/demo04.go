package main

import "fmt"

func main() {
	// 定义切片
	var slice3 []int = []int{100, 200, 300}
	//append追加元素
	slice3 = append(slice3, 400, 500, 600)
	//将切片追加到切片
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3", slice3)

	slice4 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice4)
	//操作删除2,3,4
	slice4 = append(slice4[:1], slice4[4:]...)
	fmt.Println(slice4)
}
