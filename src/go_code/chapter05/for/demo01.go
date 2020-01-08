package main

import (
	"fmt"
)

func main()  {
	for i :=1; i <= 10; i++ {
		fmt.Println("你好，我是爱冷")
	}

	//for循环的第二种写法
	j := 1
	for j <= 10 {
		fmt.Println("你好")
		j++ //循环变量迭代
	}

	/*
	//for循环的第三种写法(死循环)
	for {
		fmt.Println("ok~")
	}
	*/

	k := 1
	for { //这里也等价 for ; ; {
		if k < 10 {
			fmt.Println("ok~")
		} else {
			break //跳出整个for循环
		}
		k++
	}

	//for-range方式遍历字符串,按照字符方式，有中文也ok，注意看有中文的index
	//1个中文字符等于3个字节 (utf8)
	str := "abc~ok武汉"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}

}