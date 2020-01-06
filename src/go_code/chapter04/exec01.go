package main

import (
	"fmt"
)

func main()  {
	
	//交换变量，但不允许用中间变量
	a := 2
	b := 9

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a=", a, "b=", b)
}