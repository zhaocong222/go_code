package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	defer srcFile.Close()

	//通过srcfile,获取到Reader
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer dstFile.Close()

	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func main() {
	//拷贝图片

	//原文件
	srcFile := "d:/file/aaa.jpg"
	//目标文件
	dstFile := "d:/file/bbb.jpg"

	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("拷贝错误,err=", err)

	}
}
