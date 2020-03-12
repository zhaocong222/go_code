package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//var allChan chan interface{}
	allChan := make(chan interface{}, 5)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"大熊猫", 4}
	allChan <- cat

	//我们希望管道中第3个元素，则先推出前2个
	<-allChan
	<-allChan

	newCat := <-allChan

	fmt.Printf("newCat=%T, newCat=%v", newCat, newCat)
	//这里不能直接打印newCat.Name，因为编译的时候newCat是 interface{}类型
	//fmt.Printf("newCat.Name=%v", newCat.Name)

	//通过类型断言转换
	if v, ok := newCat.(Cat); ok {
		fmt.Printf("v.Name=%v", v.Name)
	}

}
