package main

import "fmt"

//一个被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n-1; i++ {
		res += i
	}
	return res
}

func main() {
	res := addUpper(10)
	if res != 55 {
		fmt.Println("错误")
	} else {
		fmt.Println("正确")
	}
}
