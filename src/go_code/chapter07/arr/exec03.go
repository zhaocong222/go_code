package main

import "fmt"

func main() {
	//请出一个数组的平均值
	nums := [...]int{2, 3, 4, 5}
	sum := 0

	for _, val := range nums {
		sum += val
	}

	//如何让平均值保留到小数，这里要强制转换类型， golang中没有自动转换类型，和其他语言不一样，必须强制转换
	fmt.Printf("sum=%v, 平均值是%v", sum, float64(sum)/float64(len(nums)))

}
