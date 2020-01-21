package main

import "fmt"

type MethodUtils struct {
}

func (m MethodUtils) print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//计算面积
func (m MethodUtils) mianji(len float64, width float64) float64 {
	return len * width
}

func main() {
	//1.编写结构体（MethodUtils），编写一个方法，方法不需要参数
	//在方法中打印一个10*8的矩形，在main方法重调用该方法
	var me MethodUtils
	me.print()

	fmt.Println("矩形面积是:", me.mianji(1.2, 3.2))
}
