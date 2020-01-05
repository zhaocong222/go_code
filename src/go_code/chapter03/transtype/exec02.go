package main

//如果没有使用到一个包，但是又不想删除，前面加_表忽略
import (
	_"fmt" 
)

//基本数据类型的转换.go中只有强制转换

func main() {
	/*
	var n1 int32 = 12
	var n2 int64
	var n3 int8

	//cannot use n1 + 20 (type int32) as type int64 in assignment
	//n1 + 20 还是 int32类型, int32交给n3 (int64类型) 错误，go中只能强制转换类型，所以编译不通过
	n2 = n1 + 20
	n3 = n1 + 20

	//修改
	n2 = int64(n1) + 20

	fmt.Println("n2=", n2, "n3=", n3)
	*/

	//int8 最大127
	// var n1 int32 = 12
	// var n3 int8
	// var n4 int8
	// n4 = int8(n1) + 127  //溢出
	// n3 = int8(n1) + 128  //溢出


	

}