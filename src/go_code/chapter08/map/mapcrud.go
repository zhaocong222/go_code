package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	//增加
	cities["no3"] = "重庆"
	//修改
	cities["no2"] = "香港"
	//删除
	delete(cities, "no3")

	//查找
	val, ok := cities["no1"]
	if ok {
		fmt.Println("有no1 key值为%v\n", val)
	}

	//遍历map
	for k, v := range cities {
		fmt.Printf("key为%v,val为%v\n", k, v)
	}

	//如果希望一次性删除所有key
	//1.遍历所有key，逐一删除
	//2.直接make一个新的空间，原有的不再引用，gc自动回收
	cities = make(map[string]string)

}
