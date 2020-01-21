package main

import (
	"fmt"
	"go_code/chapter09/encapsulate/model"
)

func main() {
	person := model.NewPerson("simth")
	person.SetAge(21)
	person.SetSal(5000)
	fmt.Println(person)
	fmt.Println(*person)
	//取值
	fmt.Printf("信息:名字为%v,年龄为%v,薪水为%v", person.Name, person.GetAge(), person.GetSal())
}
