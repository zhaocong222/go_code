package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	//string的基本使用
	var address string = "北京长城 110"
	fmt.Println("address=", address)

	fmt.Println("占用字节=", unsafe.Sizeof(address))

	//字符串一旦赋值了，字符串就不能修改了，在go中字符串是不可变的
	//var str = "hello"
	//str[0] = 'a' //这里就不能去修改str的内容，即go中的字符串是不可变的

	str2 := "abc\nabc"
	fmt.Println(str2)

	//使用反引号输出文本,esc下的符号
	str3 := `
	sdfsdg
	sdgfsdfg
	sdfgsdfgs
	wrwerewr
	dfghfdgh
	`
	fmt.Println(str3)
	
	//当一个拼接的操作很长是，可以分行写，但是注意，需要将+保留在上一行
	str4 := "abcd" + "efg" +
	"hij" + "opq" +
	"uvw" + "xyz"
	fmt.Println(str4)

	//基础数据类型默认值
	var a int  //  0
	var b float32 // 0
	var c float64 // 0
	var isMarry bool // false
	var name string // ""
	// %v 表示按照变量的值输出
	fmt.Printf("a=%d,b=%f,c=%f,isMarry=%v, name=%v", a, b, c, isMarry, name)


}