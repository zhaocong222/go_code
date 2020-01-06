package main

import (
	"fmt"
)

func main()  {
	//打印该月份所属的季节， 3，4，5春。。。
	var month byte
	fmt.Println("请输入月份")
	fmt.Scanf("%d", &month)

	switch month {
		case 3, 4, 5 :
			fmt.Println("春")
		case 6, 7, 8 :
			fmt.Println("夏")
		case 9, 10, 11 :
			fmt.Println("秋")
		case 12, 1, 2 :
			fmt.Println("冬")
		default:
			fmt.Println("输入有误")
	}

}