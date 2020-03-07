package main

import "fmt"

type Usb interface {
	Say()
}

type Stu struct {
}

//*Stu指针类型
func (this *Stu) Say() {
	fmt.Println("Say()")
}

func main() {
	var stu Stu = Stu{}
	//错误! 会报Stu类型没有实现usb接口
	var u Usb = stu
	//如果希望编译通过，修改为 var u Usb = &stu
	u.Say()
}
