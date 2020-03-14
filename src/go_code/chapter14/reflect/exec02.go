package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

//结构体方法
func (s Monster) Print() {
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("----end----")
}

func (s Monster) GetNum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(s interface{}) {
	//获取reflect.Type 类型
	typ := reflect.TypeOf(s)
	//获取reflect.Valuele类型
	val := reflect.ValueOf(s)
	//获取到s对应的类型
	kd := val.Kind()
	//如果传入的不是结构体，直接退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	//遍历结构体所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:值为=$v\n", i, val.Field(i))
		//获取到struct标签，注意需要通过reflect.Type来获取tag标签
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段有tag标签就显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	//能获取结构体有多少方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//****函数的排序 是以函数的名字排序
	val.Method(1).Call(nil) //调用第2个方法 Print

	//调用结构体的第一个方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	//Call传的是 Value的切片类型, 返回的也是Value的切片类型
	res := val.Method(0).Call(params) //传入的参数是 []relect.Value

	fmt.Println("res=", res[0].Int())

}

//使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
func main() {
	a := Monster{
		Name:  "喷火龙",
		Age:   400,
		Score: 30.8,
	}
	//将Monster实例传递
	TestStruct(a)
}
