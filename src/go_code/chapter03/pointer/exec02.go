package main 

import (
	"fmt"
)

//指针类型练习
func main() {

	num := 9
	//打印num地址
	fmt.Println("num地址是\n", &num)
	
	//通过指针修改值
	var ptr *int = &num
	*ptr = 100
	fmt.Println("num=", num)


}