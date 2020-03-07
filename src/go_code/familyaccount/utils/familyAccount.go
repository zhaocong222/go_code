package utils

import "fmt"

type FamilyAccount struct {

	//声明一个变量，保存接收用户输入的选项
	key string
	//声明一个变量，控制会否退出for循环
	loop bool
	//定义账户的余额[]
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义个变量，记录是否有收支的行为
	flag bool
	//收支的详情使用字符串来记录
	//当有收支时，只需要对details进行拼接处理即可
	details string
}

//编写工厂模式方法，返回一个*FamilyAccount实例
func NewFamily() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户余额\t收支余额\t说  明",
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("-----------当前收支明细记录------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细...来一把吧")
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money //修改账户余额
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量
	//收入 11000 1000 有人发红包
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Println("登记支出.")
	fmt.Scanln(&this.money)
	//这里需要做一个必要的判断
	//支出的钱必须小于等于余额
	if this.money > this.balance {
		fmt.Println("余额不足")
		//break
	} else {
		this.balance -= this.money
		fmt.Println("本次支出说明：")
		this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
		this.flag = true
	}

}

func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗?y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}

	//yes,退出
	if choice == "y" {
		this.loop = false
	}
}

//给上面定义的结构体，绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu() {

	//显示菜单
	for { //这里也等价 for ; ; {
		fmt.Println("----------家庭收支记账软件---------")
		fmt.Println("           1 收支明细   ")
		fmt.Println("           2 登记收入   ")
		fmt.Println("           3 登记支出   ")
		fmt.Println("           4 退出软件   ")

		fmt.Print("请选择(1-4): ")

		//等待用户输入
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项..")
		}

		//退出软件，跳出循环
		fmt.Println(this.loop)
		if !this.loop {
			break
		}

	}
}
