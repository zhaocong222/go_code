package main

import "fmt"

//基本数据类型的转换.go中只有强制转换

func main() {
	i := 100
	//希望将 i => float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)

	fmt.Printf("i=%v, n1=%v, n2=%v, n3=%v", i, n1, n2, n3)

	///注意：被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化

	fmt.Printf("\n i数据类型%T", i)

}