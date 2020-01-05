package main

import (
	"fmt"
	//使用model/utils.go 文件里面的变量，需要先引入该包
	//绝对路径为 D:\goproject\src\go_code\chapter03\demo\model ,这里设置过GOPATH，默认会加载GOPATH/src
	"go_code/chapter03/demo/model"
)

func main()  {
	
	//使用utils.go的heroName变量
	fmt.Println(model.HeroName)

}