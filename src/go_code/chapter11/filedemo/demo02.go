package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("d:/file/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时的关闭file
	defer file.Close() //函数结束后，关闭文件句柄

	//创建一个 *Reader, 带缓冲区, 默认缓冲区为4096个字节
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行结束一次
		if err == io.EOF {                  //表示文件的末尾
			break
		}

		//输出内容
		fmt.Print(str)
	}

	fmt.Println("文件读取结束....")
}
