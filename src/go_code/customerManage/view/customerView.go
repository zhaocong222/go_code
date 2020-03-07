package main

import (
	"fmt"
	"go_code/customerManage/model"
	"go_code/customerManage/service"
)

//不用大写，就在本文件用
type customerView struct {
	key             string //接收用户的输入....
	loop            bool   //表示是否循环显示菜单
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list() {
	//首先，获取到当前所有的客户信息(切片中)
	customers := this.customerService.List()
	//显示
	fmt.Println("-------------客户列表---------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------- 客户列表完成 ---------------")

}

//得到用户的输入信息，构建新的客户，并完成添加
func (this *customerView) addCustomer() {
	fmt.Println("-------添加客户-------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	//id需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("-------添加成功-------")
	} else {
		fmt.Println("-------添加失败-------")
	}
}

//删除用户
func (this *customerView) delete() {
	fmt.Println("--------删除客户---------")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}

	fmt.Println("确认是否删除(Y/N):")
	choice := ""
	fmt.Scanln(&choice)

	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("------删除完成-----")
		} else {
			fmt.Println("------删除失败，输入id不存在-----")
		}
	}
}

//退出
func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&this.key)
		if this.key == "n" || this.key == "y" {
			break
		}

		fmt.Println("你的输入有误，请重新输入y/n")
	}

	if this.key == "y" {
		this.loop = false
	}
}

//显示主菜单
func (this *customerView) mainMenu() {

	for { //这里也等价 for ; ; {
		fmt.Println("----------客户信息管理软件---------")
		fmt.Println("           1 添加客户   ")
		fmt.Println("           2 修改客户   ")
		fmt.Println("           3 删除客户   ")
		fmt.Println("           4 客户列表   ")
		fmt.Println("           5 退    出   ")

		fmt.Print("请选择(1-5): ")

		//等待用户输入
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.addCustomer()
		case "2":

		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("请输入正确的选项..")
		}

		//退出软件，跳出循环
		fmt.Println(this.loop)

		if !this.loop {
			break //退出
		}

	}

	fmt.Println("你退出了客户关系系统")

}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}

	customerView.customerService = service.NewCustomerService()

	customerView.mainMenu()
}
