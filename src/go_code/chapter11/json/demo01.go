package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体 tag标签： `json:"name"` 序列化后的名字
type Monster struct {
	Name     string `json:"name"` //反射机制
	Age      int    `json:"monster_age"`
	Birthday string
	Sal      float64
	Skill    string
}

//对结构体序列化
func testStruct() {
	monster := Monster{
		Name:     "悟空",
		Age:      1000,
		Birthday: "1000-11-21",
		Sal:      30000,
		Skill:    "叫人",
	}

	//序列化,返回[]byte切片，具体看官网文档
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("monster序列化后=%v\n", string(data))
}

//将map序列化
func testMap() {
	//key为string类型，值为任意类型
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	//map默认是引用类型，就不用传指针了
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("a序列化后=%v\n", string(data))

}

//对切片进行序列化
func testSlice() {
	//切片里有多个map
	var slice []map[string]interface{}
	map1 := make(map[string]interface{})
	map1["name"] = "jack"
	map1["age"] = 7
	map1["address"] = "北京"

	slice = append(slice, map1)

	map2 := make(map[string]interface{})
	map2["name"] = "tome"
	map2["age"] = 17
	map2["address"] = [2]string{"日本", "夏威夷"}

	slice = append(slice, map2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice序列化后=%v\n", string(data))

}

//基本数据类型序列化,意义不大
func testFloat64() {
	num1 := 2345.67
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("num1序列化后=%v\n", string(data))
}

func main() {
	//演示将结构体，map，切片进行序列化
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
