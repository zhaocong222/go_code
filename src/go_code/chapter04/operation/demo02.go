package main

import (
	"fmt"
)

func main()  {
	//在golang中，++和-- 只能独立使用
	var i int = 8
	var a int
	a = i++ //错误，i++只能独立使用
	a = i-- //错误，i--只能独立使用

}