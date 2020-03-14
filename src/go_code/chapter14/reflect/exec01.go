package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		str := "tom"
		fs := reflect.ValueOf(str) //ok fs-> string类型
		fs.SetString("jack")       //error 指针类型才能调用SetString
	*/

	//正确代码片段
	str := "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jakc")
	fmt.Println(str) //jakc

}
