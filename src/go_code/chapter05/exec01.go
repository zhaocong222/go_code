package main

import (
	"fmt"
)

func main()  {
	
	//编写程序，声明2个int32型变量并赋值，判断两数之和，如果大于等于50,打印"hello world"
	var a,b int32 = 12,40
	if a + b >= 50 {
		fmt.Println("hello world")
	}

	//编写程序，声明2个float64型变量并赋值，判断第一个数大于10.0
	//且第2个数小于20.0,打印两数之和
	var n1,n2 float64 = 11.2,15.3
	if n1 > 10.0 && n2 < 20.0 {
		fmt.Println("n1+n2=", n1+n2)
	}

	//定义两个变量int32, 判断二者的和，是否能被3又能被5整除，打印提示信息
	var n3, n4 int32 = 3, 12
	if (n3 + n4) % 3 == 0 && (n3 + n4) % 5 == 0 {		
		fmt.Println("haha")
	}  

	
}