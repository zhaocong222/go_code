package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil.ReadFile 一次性将文件读取到位,只适合文件不太大的情况下
	file := "d:/file/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}

	//把读取到的内容显示到终端
	//没有显式的open文件，因此也不需要显示的close文件
	//fmt.Printf("%v", content) //[]byte 输出会按照数字打印
	fmt.Printf("%v", string(content))
}
