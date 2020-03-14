package main

import (
	"fmt"
	"reflect"
)

//通过反射
//修改 num int的值
//修改 student的值
func reflectTest03(b interface{}) {
	//获取reflect.value
	rVal := reflect.ValueOf(b)
	//查看 rVal的kind, 传入的是地址，此刻为指针类型
	fmt.Println("rVal kind=%v", rVal.Kind())
	//改变值,通过Elem获取指针的值
	rVal.Elem().SetInt(100)

}

func main() {
	var num int = 10
	//传入地址
	reflectTest03(&num)

	fmt.Println("num=", num) //100

}
