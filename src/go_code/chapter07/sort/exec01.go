package main

import "fmt"

//冒泡排序
func BubbleSort(arr *[5]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}
}

func main() {
	// 1. 一共经过 arr.len - 1次的轮数比较，每一轮将会确定一个数的位置
	// 2. 每一轮的比较次数再逐渐的减少
	// 3. 当发现前面的一个数比后面的一个数大的时候，就进行交换
	arr := [...]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
	//打印数组
	fmt.Println(arr)
}
