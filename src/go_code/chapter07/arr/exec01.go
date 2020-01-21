package main

import "fmt"

func main() {
	//1.创建一个byte类型的26个元素的数组，分别放置'A-Z'
	mychars := [26]byte{}

	for i := 0; i < 26; i++ {
		mychars[i] = 'A' + byte(i)
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", mychars[i])
	}

}
