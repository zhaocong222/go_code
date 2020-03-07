package main

import "fmt"

//声明一个接口,实现接口的类必须实现所有方法
type Usb interface {
	//2个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

//让phone实现usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
}

//让Camera实现 usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

//计算机
type Computer struct {
}

//编写一个方法working方法，接收一个usb接口类型变量
//只要是实现了usb接口，所谓实现usb接口，就是指实现了usb接口声明所有的方法
func (c Computer) Working(u Usb) {
	u.Start()
	u.Stop()
}

//让camera实现usb接口的方法
func main() {
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	//关键点
	computer.Working(phone)
	computer.Working(camera)

}
