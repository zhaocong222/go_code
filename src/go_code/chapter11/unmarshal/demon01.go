package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"` //反射机制
	Age      int    `json:"monster_age"`
	Birthday string
	Sal      float64
	Skill    string
}

func unmarshalStruct() {
	str := "{\"name\":\"悟空\",\"monster_age\":1000,\"Birthday\":\"1000-11-21\",\"Sal\":30000,\"Skill\":\"叫人\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v", monster)
}

//反序列化
func main() {
	unmarshalStruct()
}
