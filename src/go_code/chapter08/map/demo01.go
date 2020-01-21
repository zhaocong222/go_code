package main

import "fmt"

func main() {
	//map的声明和注意事项, key为字符串型，值为字符串型
	var a map[string]string
	//在使用map前，需要先make,make的作用就是给map分配数据空间
	//这里分配10个空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "无用"
	a["no1"] = "武松"
	a["no3"] = "吴用"

	fmt.Println(a)

}
