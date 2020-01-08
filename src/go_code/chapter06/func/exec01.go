package main

import "fmt"

func peach(day int) int {

	if day > 10 || day < 1 {
		fmt.Println("输入的天数不对")
		return 0
	}

	if day == 10 {
		return 1
	} else {
		return (peach(day+1) + 1) * 2
	}
}

//第10天只有1个桃子
//第9天有几个桃子 = (第10天的桃子 + 1) * 2 = 4
//第n天的桃子 =  (( n + 1 )天的桃子 + 1) * 2
func main() {

	var day int = 9
	total := peach(day)

	fmt.Println("最初一共有桃子:", total)
}
