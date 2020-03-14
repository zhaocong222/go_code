package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest01(b interface{}) {
	//通过反射获取传入的变量的type,kind,值
	//1.先获取到 reflect.type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	//rVal 此时rVal不是int类型，运算需要反射的方法
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)

	//获取此刻真实的类型
	fmt.Printf("rVal=%v rval type=%T\n", rVal, rVal)

	//下面我们将rVal转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转换成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

func main() {
	//请编写一个案例
	//演示对(基本数据类型、interface{}、relect.Value进行反射的基本操作)
	var num int = 100
	reflectTest01(num)
}
