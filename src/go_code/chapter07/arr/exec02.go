package main

import "fmt"

func main() {
	//求出一个数组的最大值，获取对应下标
	nums := [...]int{234, 345, 123, 45, 64764, 21, 43, 56}

	maxVal := nums[0]
	maxIndex := 0

	for index := 0; index < len(nums); index++ {
		if maxVal < nums[index] {
			maxVal = nums[index]
			maxIndex = index
		}
	}

	fmt.Printf("最大值是:%v,下表是:%v", maxVal, maxIndex)
}
