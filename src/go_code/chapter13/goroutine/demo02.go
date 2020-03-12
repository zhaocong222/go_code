package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNUm := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNUm)

	//设置num - 1的 cpu运行go程序, 1.9以后默认
	runtime.GOMAXPROCS(cpuNUm - 1)
	fmt.Println("ok")
}
