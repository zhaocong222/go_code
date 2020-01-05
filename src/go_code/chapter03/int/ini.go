package main 

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 1
	fmt.Println("i=", i)

	//测试一下 int8范围
	//其他的 int16 int32 int64累推
	//var j int8 = 128 //constant 128 overflows int8
	//fmt.Println("j=", j)
	
	//测试一下无符号 (uint8 0 -255)
	//var k uint8 = 256
	//fmt.Println("k=", k)

	//int, uint, rune, byte的使用
	var a int = 8900
	fmt.Println("a=", a)

	//整形的使用细节
	var n1 = 100 //? n1 是什么类型
	//这里我们查看某个变量数据类型
	//%T 类型, 默认是int， 32系统就是int32,64就是int64
	fmt.Printf("n1 的 类型 %T", n1)

	//如何在程序查看某个变量的占用字节大小和数据类型
	var n2 int64 = 10
	//unsafe.Sizeof(n1) 是unsafe包的一个函数，可以返回n1变量占用的字节数
	fmt.Printf("n2的类型%T n2占用的字节数是 %d", n2, unsafe.Sizeof(n2))

	//golang程序中整型变量在使用时，遵守保小不保大的原则 
	//即：在保证程序正确运行下，尽量使用占用空间小的数据类型
	
	
}