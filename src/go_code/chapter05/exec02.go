package main

import (
	"fmt"
)

func main()  {
	
	var score int
	fmt.Println("请输入成绩:")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("奖励一辆bmw")
	} else if score > 80 && score <=99 {
		fmt.Println("奖励一台iphone7plus")
	} else if score >= 60 && score <=80 {
		fmt.Println("奖励一个ipad")
	} else {
		fmt.Println("什么都不奖励")
	}

}