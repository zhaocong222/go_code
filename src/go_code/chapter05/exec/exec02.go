package main

import (
	"fmt"
)


func main()  {

	var year int
	fmt.Println("请输入年份")
	fmt.Scanln(&year)

	if year % 4 == 0 {
		fmt.Printf("%v是闰年", year)
	}

}