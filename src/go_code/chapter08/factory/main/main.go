package main

import (
	"fmt"
	"go_code/chapter08/factory/model"
)

func main() {
	//创建要给student实例
	// stu := model.Student{
	// 	Name:  "tom",
	// 	Score: 78.9,
	// }
	//fmt.Println(stu)

	//注意包里结构体的字段大小写
	stu1 := model.NewStudent("jack", 25) //返回的指针
	fmt.Println(stu1)                    // &{...}
	fmt.Println(*stu1)                   //{jack 21}
	//stu1.Name 底层会自动优化成(*stu1).Name, 当然可以直接写成(*stu1).Name
	fmt.Printf("name=%v,score=%v\n", stu1.Name, stu1.GetScore())
	fmt.Printf("name=%v,score=%v", (*stu1).Name, (*stu1).GetScore())
}
