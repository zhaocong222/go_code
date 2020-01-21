package main

import "fmt"

func main() {
	//第一种使用方式
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

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天井"
	cities["no3"] = "上海"

	//第三种方式
	/*
		var heroes map[string]string = map[string]string{
			"hero1": "宋江",
			"hero2": "卢俊义",
		}
	*/
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
	}

	fmt.Println(heroes)

	//案例演示1. 存放3个学生的信息，每个学生有name和sex信息
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["address"] = "北京长安街"
	studentMap["stu01"]["sex"] = "男"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "tom"
	studentMap["stu02"]["address"] = "上海"
	studentMap["stu01"]["sex"] = "女"

}
