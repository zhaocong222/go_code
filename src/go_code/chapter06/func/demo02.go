package main

import (
	"fmt"
)

func getSumAndSub(n1 int, n2 int) (int, int) {

	res1 := n1 + n2
	res2 := n1 - n2
	return res1, res2

}

func main() {

	res1, res2 := getSumAndSub(10, 20)
	fmt.Printf("res1=%v res2=%v\n", res1, res2)

}
