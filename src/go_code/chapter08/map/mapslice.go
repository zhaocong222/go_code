package main

import "fmt"

func main() {
	//演示map切片的使用
	var monsters []map[string]string
	//切片make,1个空间
	monsters = make([]map[string]string, 1)
	//增加第一个妖怪信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	//使用切片的append动态增加monsters
	newMonster := map[string]string{
		"name": "玉兔精",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}
