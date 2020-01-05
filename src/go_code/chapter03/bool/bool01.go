package main

import "fmt"
import "unsafe"

func main()  {
	b := false
	fmt.Println("b=", b)

	//占用空间
	fmt.Println("b的占用空间=", unsafe.Sizeof(b))
}