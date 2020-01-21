package main

import "fmt"

func sum(n1 int, n2 int) int {

	//当执行到defer时，会将defer后面的语句压入到独立的栈(defer栈)
	//当整个函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	//将n1压入到栈中，是把n1的值拷贝入栈, 注意不是引用
	defer fmt.Println("ok1 n1=", n1) //defer  4.执行  n1 = 10
	defer fmt.Println("ok2 n2=", n2) //defer  3.执行  n2 = 20

	n1++           //n1 = 11
	n2++           //n2 = 21
	res := n1 + n2 // 32

	fmt.Println("ok3 res=", res) //1.先执行
	return res                   //2. 再执行，函数执行完毕，开始从defer栈出栈
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res) //5.最后执行
}
