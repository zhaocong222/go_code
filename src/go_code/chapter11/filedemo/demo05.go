package main

import (
	"fmt"
	"io/ioutil"
)

//将一个文件里的内容，写入到另一个文件中
func main() {
	//将 d:/file/abc.txt 文件写入到 d:/file/demo.txt

	file1Path := "d:/file/abc.txt"
	file2Path := "d:/file/demo.txt"

	content, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("读取文件出错")
		return
	}

	err = ioutil.WriteFile(file2Path, content, 0666)
	if err != nil {
		fmt.Println("写文件出错err=%v", err)
		return
	}
}
