package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	//概念说明:file的叫法
	file, err := os.Open("d:/file/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//输出下文件，看看文件是什么, file就是一个指针 *File
	fmt.Printf("file=%v", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
