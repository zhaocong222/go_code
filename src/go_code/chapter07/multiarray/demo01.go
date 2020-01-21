package main

import "fmt"

func main() {
	//二位数组定义
	var arr [2][2]int //int64 8字节
	arr[1][1] = 1

	fmt.Println(arr)
	//arr[0]的地址0xc000012480 arr[1]的地址0xc000012490 =>相差了16个字节
	//因为一维数组里含有2个元素，int64->8字节  2 * 8 = 16
	fmt.Printf("arr[0]的地址%p\n", &arr[0])
	fmt.Printf("arr[1]的地址%p\n", &arr[1])

	//数组的地址，即头地址
	fmt.Printf("arr[0][0]的地址%p\n", &arr[0][0])
	fmt.Printf("arr[1][0]的地址%p\n", &arr[1][0])

	//直接初始化
	//var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr3)
}
