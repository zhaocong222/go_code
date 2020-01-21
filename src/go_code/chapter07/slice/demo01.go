package main

import "fmt"

func main() {
	intArr := [...]int{1, 22, 33, 66, 99}
	//切片
	slice := intArr[1:3]                    //引用下标为1的元素到下表为3的元素(不包含3)
	fmt.Println("intArr=", intArr)          //[1, 22, 33, 66, 99]
	fmt.Println("slice的元素是=", slice)        //[22, 33]
	fmt.Println("slice的元素个数是=", len(slice)) //2
	fmt.Println("slice的容量=", cap(slice))

	//slice是引用类型
	//从底层来说slice的第一个元素引用的就是intArr下标为1的元素地址
	//打印地址，验证
	fmt.Printf("slice的一个元素地址是%p,intArr下标为1的元素地址%p",
		&slice[0], &intArr[1])

	//修改slice第一个元素，则intArr会发生改变，因为引用
	slice[0] = 777
	fmt.Println("修改后的slice", slice)
	fmt.Println("修改slice后intArr的变化", intArr)

}
