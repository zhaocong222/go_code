package main

import "fmt"

//n1 n2 都为float32类型
func sum(n1, n2 float32) float32 {
	fmt.Printf("n1 type = %T\n", n1)
	return n1 + n2
}

//编写一个函数交换n1和n2的值
func swap(n1, n2 *int) {
	*n1 = *n1 + *n2
	*n2 = *n1 - *n2
	*n1 = *n1 - *n2
}

func main() {
	fmt.Println("sum=", sum(1, 2)) //3

	//调用
	var a, b int
	a = 1
	b = 2
	fmt.Printf("a=%v, b=%v\n", a, b)

	//调用a,b值
	swap(&a, &b)

	fmt.Printf("a=%v, b=%v\n", a, b)

}
