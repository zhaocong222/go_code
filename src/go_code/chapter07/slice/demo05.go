package main

import "fmt"

func main() {
	//字符串底层就是byte数组
	str := "hello@atguigu"
	//使用切片
	slice := str[6:]
	fmt.Println("slice=", slice)

	//如果需要修改字符换，可以先转换
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

	//细节,我们转成[]byte后，可以处理英文和数字，不能处理中文
	//原因[]byte字节来处理，而一个汉字，是3个字节，因此就会出现乱码
	//转换为[]rune即可，rune是按字符处理的
	arr2 := []rune(str)
	arr2[0] = '哈'
	str = string(arr2)
	fmt.Println("str=", str)
}
