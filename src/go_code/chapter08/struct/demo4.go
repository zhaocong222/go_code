package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Monster struct {
	Name  string `json:"name"` //`json:"name"` 就是 struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	a = A(b) //这里结构体转行，必须满足结构体字段完全一样
	fmt.Println(a, b)

	//1. 创建一个Monster变量
	monster := Monster{"牛魔王", 400, "芭蕉扇"}
	//2. 将monster变量序列化为 json格式字串
	//这个jsonStr是[]byte byte切片
	//函数底层利用反射
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处错误", err)
	}
	fmt.Println("jsonstr", jsonStr)

	fmt.Println("json字串", string(jsonStr))
}
