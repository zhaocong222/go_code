package main

import (
	"flag"
	"fmt"
)

func main() {

	//定义几个变量，用于接收命令行输入的参数
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	// "u", 就是 -u 指定参数
	// "" 默认值
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "host", "localhost", "主机名")
	flag.IntVar(&port, "port", 3306, "端口")

	//转换，必须调用此方法
	flag.Parse()

	//输出结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v",
		user, pwd, host, port)
}
