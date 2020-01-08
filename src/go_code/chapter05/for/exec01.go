package main

import (
	"fmt"
)

func main()  {
	
	var num, sum int = 0, 0

	for i :=1; i <= 100; i++ {
		if i % 9 == 0 {
			num++
			sum += i
		}

	}

	fmt.Printf("个数是%d,总和是%d", num, sum)


	var n int = 6
	for h := 0; h <= 6; h++ {
		fmt.Printf("%v + %v = %v\n", h, n - h, n)
	}
}