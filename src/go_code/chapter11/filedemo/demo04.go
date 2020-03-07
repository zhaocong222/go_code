package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "d:/file/abc.txt"
	//O_WRONLY只写模式,O_CREATE 如果文件不存在就创建
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	//准备写入5句 "hello,Gardon"
	str := "hello,Gardon\r\n"
	//写入时，使用带缓冲区的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为write是带缓存的，此时文件内容还在缓冲区中，需要将缓冲区的内容刷入到文件中
	err = writer.Flush()
	if err != nil {
		fmt.Printf("刷入文件失败")
	}
}
