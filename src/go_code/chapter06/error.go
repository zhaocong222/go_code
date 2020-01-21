package main

import (
	"errors"
	"fmt"
)

func test() {
	//使用defer + recover来捕获和处理异常
	//将此匿名函数压入栈中，函数执行完后出栈
	defer func() {
		err := recover() //内置函数，可以捕获到异常
		if err != nil {  //说明捕获到错误
			fmt.Println("捕获到了错误:", err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//函数去读取配置文件init.conf的信息
//如果文件名不正确，我们就返回一个自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误。。。")
	}
}

func test02() {
	err := readConf("config2.ini")
	if err != nil {
		panic(err)
	}

	fmt.Println("继续执行")

}

func main() {
	//test()
	test02()
	fmt.Println("下面的代码")
}
