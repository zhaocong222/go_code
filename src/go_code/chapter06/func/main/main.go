package main

//重点：在import包时，路径从$GOPATH的src下开始，不用带src,编译器会自动从src下开始引入
import (
	"fmt"
	//"go_code/chapter06/func/utils"
	//别名方式
	haha "go_code/chapter06/func/utils"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 3.3
	var operator byte = '*'
	//这里的utils是包名，不是文件名，utils.go文件里package后的命名
	//result := utils.Cal(n1, n2, operator)
	//利用别名
	result := haha.Cal(n1, n2, operator)
	fmt.Println("result=", result)
}
