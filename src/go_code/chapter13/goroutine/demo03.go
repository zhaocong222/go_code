package main

import (
	"fmt"
	"sync"
	"time"
)

//需求：现在要计算 1-200的各个数的阶乘，并把各个数的阶乘放入到map中
//最后显示出来，要求使用goroutine完成

//思路
//1.编写一个函数，来计算各个数的阶乘，并放入到map中
//2.我们启动的协程多个，统计的将结果放入到map中
//3.map应该做出一个全局的

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁
	//sync是包
	lock sync.Mutex
)

//计算n的阶乘,结果放入map
func test2(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//加锁 ,不同几个协程 同时刻往map里面写，需要加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {
	for i := 1; i <= 10; i++ {
		//启动200个协程
		go test2(i)
	}

	// 休眠5秒, 到底等多少？后面解决
	time.Sleep(3 * time.Second)

	//这里我们输出结果，这里也要加锁，可能底层并不知道协程写数据完了，已经解锁
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
	lock.Unlock()
}
