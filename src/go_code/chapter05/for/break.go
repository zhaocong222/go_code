package main

import (
	"fmt"
)

func main()  {
	
	lable2: //设置标签
	for i := 0; i < 4; i++ {
		//lable1: //设置标签
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable2 //break 默认会跳出最近的for循环
			}
			fmt.Println("j=", j)
		}
	}	
}