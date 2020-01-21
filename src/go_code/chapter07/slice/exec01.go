package main

import "fmt"

func fbc(n int) []int64 {
	//声明有n个元素的切片
	fbnslice := make([]int64, n)
	//第一个数和第二个数的斐波那契为1,1
	fbnslice[0], fbnslice[1] = 1, 1
	for index := 2; index < n; index++ {
		fbnslice[index] = fbnslice[index-1] + fbnslice[index-2]
	}

	return fbnslice
}

func main() {
	/**
	 * 1.可以接受一个 nindexint
	 *  2.能够将斐波那契的数	fbnslice列放到切片中
	 *  3.提示，arr[0] = 1; arr[1] = 1; ....
	 */
	slice := fbc(6)
	fmt.Println(slice)
}
