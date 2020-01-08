package main

import (
	"fmt"
)

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//这时，myFun 就是 func (int, int) int 类型
type myFunType func(int, int) int

func myFun2(funx myFunType, num1 int, num2 int) int {
	return funx(num1, num2)
}

//支持对函数返回值命名,定义了返回的参数
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	a := getSum
	fmt.Println("a的类型%T, getSum类型是%T\n", a, getSum)

	res := a(10, 40)
	fmt.Println("res=", res)

	//看案例
	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)

	type myInt int //给int取了别名,在go中myInt和int虽然都是int类型，但是go认为两个是不同类型
	var num1 myInt
	num1 = 40
	num2 := int(num1) //各位，注意这里依然需要显示转换，go认为myInt和int两个类型
	fmt.Println("num1=", num1, "num2=", num2)

	//再加一个案例
	res3 := myFun2(getSum, 1, 11)
	fmt.Println("res3=", res3)

	//支持对函数返回值命名
	sum, sub := getSumAndSub(1, 2)
	fmt.Println(sum, sub)
}
