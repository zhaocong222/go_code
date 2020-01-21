package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机生成5个数，并将其反转打印
	//1.随机生成5个数，放入数组中
	//2.反转打印
	var intArr3 [5]int
	//随机种子 ， time.Now().Unix()当前时间戳秒, UnixNano纳秒
	rand.Seed(time.Now().UnixNano())

	for index := 0; index < 5; index++ {
		intArr3[index] = rand.Intn(100) // <= 100
	}

	fmt.Println(intArr3)

	//2.反转
	for i := 0; i < len(intArr3)/2; i++ {
		temp := intArr3[len(intArr3)-1-i]
		intArr3[len(intArr3)-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println(intArr3)
}
