package main

import "fmt" //fmt包中提供格式化，输出，输入的函数

//用gofmt -w main.go 格式化文件空格缩进

func main() {
	// 演示转义字符的使用 \t
	fmt.Println("tom\tjack")
	// 换行
	fmt.Println("hello\nworld")
	// 转义斜杠
	fmt.Println("c:\\User\\lemon")
	// 转义双引号
	fmt.Println("tom说\"i love you\"")
	// \r　回车 ，从当前行的最前面开始输出，覆盖掉以前的内容
	fmt.Println("天龙八部飞山飞狐\r张飞")

}
