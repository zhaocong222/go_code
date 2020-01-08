package main

import (
	"fmt"
)

func main()  {
	var num int
	fmt.Println("请输入一个整数")
	fmt.Scanln(&num)

	if num > 0 {
		fmt.Println("大于0")
	} else if num < 0 {
		fmt.Println("小于0")
	} else {
		fmt.Println("等于0")
	}

}