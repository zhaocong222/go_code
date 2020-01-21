package main

import "fmt"

func main() {
	var age [3]int //int64占8个字节,int32占4个字节
	age[0] = 12
	age[1] = 43

	for index, val := range age {
		fmt.Println(index, val)
	}
	//数组的地址就是数组第一个元素的地址
	//数组元素的地址是连续的, 第二个元素的地址 就是 在第一个元素的地址上 加上类型的字节数,以此类推
	fmt.Printf("age数组的地址%p, age[0]的地址是%p, age[1]的地址是%p",
		&age, &age[0], &age[1])

}
