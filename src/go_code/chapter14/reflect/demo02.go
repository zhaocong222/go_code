package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {

	rTyp := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)

	//两种方式获取变量对应的Kind
	kind1 := rVal.Kind()
	kind2 := rTyp.Kind()
	fmt.Printf("kind=%v kind=%v\n", kind1, kind2)

	iV := rVal.Interface()

	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

func main() {
	//请编写一个案例
	//演示对(基本数据类型、interface{}、relect.Value进行反射的基本操作)
	stu := Student{"悟空", 20}
	reflectTest02(stu)
}
