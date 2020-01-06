package main

import (
	"fmt"
)

func main()  {
	
	//练习1. 距离放假还有97天，还有多少星期多少天
	days := 97
	var week,day int
	week = days / 7
	day = days % 7

	fmt.Printf("距离放假还有%d星期，%d天\n", week, day)

	//定于一个变量保存华氏温度，华氏温度转换摄氏度的公式为：
	// 5/9*（华氏温度 - 100）,请求出华氏温度对应的摄氏度
	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("对应的摄氏度是%f", sheshi)
}