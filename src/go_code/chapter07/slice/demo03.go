package main

import "fmt"

func main() {
	//使用常规的for循环遍历切片
	arr := [...]int{10, 20, 30, 40, 50} //定义数组
	slice := arr[1:4]                   //[20,30,40]

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}

	fmt.Println()

	//使用for-range
	for i, val := range slice {
		fmt.Printf("slice[%v]=%v ", i, val)
	}

	slice2 := slice[1:2] //[30]
	fmt.Println("\nslice2=", slice2)

	slice2[0] = 100
	fmt.Println("arr=", arr) //[10 20 100 40 50]
}
