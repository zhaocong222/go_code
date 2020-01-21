package main

import "fmt"

//数组传递是值类型,传递数组必须要写长度
//[3]int 数组,  []int切片
func test01(arr [3]int) {
	arr[0] = 88
}

func test02(arr *[3]int) {

	//打印地址
	fmt.Printf("地址传递，函数内打印地址:%p\n", arr)

	//arr是指针，所以(*arr)才是地址的值
	(*arr)[0] = 99
}

func main() {
	//数组是值类型
	arr := [...]int{11, 22, 33}
	test01(arr)
	fmt.Println(arr)

	//查询数组地址
	fmt.Printf("地址:%p\n", &arr)
	//传地址改变值
	test02(&arr)
	fmt.Println(arr)

}
