package main

import (
	"fmt"
)


func main()  {
	
	//使用switch把小写类型的char型转为大写（键盘输入）
	//只转换a,b,c,d,e.其他的输出"other"
	var char byte
	fmt.Println("请输入小写字母")
	fmt.Scanf("%c", &char)

	switch char {
		case 'a', 'b', 'c', 'd':
			char -= 32
			fmt.Printf("%c", char) 
		default:
			fmt.Println("other")
	}


}