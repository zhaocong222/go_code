package main

import (
	"fmt"
)

//传入地址
func test03(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test03() n1=", *n1) //30
}

func main() {
	num := 20
	test03(&num)
	fmt.Println("main () num=", num) //30
}
