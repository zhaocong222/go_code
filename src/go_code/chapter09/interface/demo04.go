package main

import "fmt"

type Stu struct{}

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool类型，值是%v\n", index+1, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 整数 类型，值是%v\n", index+1, x)
		case string:
			fmt.Printf("第%v个参数是 字符串 类型，值是%v\n", index+1, x)
		case Stu:
			fmt.Printf("第%v个参数是 Stu 类型，值是%v\n", index+1, x)
		case *Stu:
			fmt.Printf("第%v个参数是 *Stu 类型，值是%v\n", index+1, x)
		}
	}
}

func main() {
	a := 1
	c := true
	d := "hello"

	stu1 := Stu{}
	stu2 := &Stu{} //指针类型

	TypeJudge(a, c, d, stu1, stu2)

}
