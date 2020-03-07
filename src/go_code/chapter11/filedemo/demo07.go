package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//定义一个结构体，用于保存统计的结果
type CharCount struct {
	ChCount    int //记录英文个数
	NumCount   int //记录数字个数
	SpaceCount int //记录空格的个数
	OtherCount int //记录其他字符的个数
}

func main() {
	//统计一个文件中含有的英文，数字...
	//思路,打开一个文件，创建一个reader
	//每读取一行，就统计该行有多少个...
	//然后将结果保存到一个结构体
	fileName := "d:/file/demo07.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	//初始化结构体
	var count CharCount
	//创建带缓冲区的Reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //读到文件末尾
			break
		}

		//遍历str,进行统计
		for _, v := range str {
			//fmt.Println(v)
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透处理
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}

	}

	//输出统计的结果看看是否正确
	fmt.Println(count)
}
