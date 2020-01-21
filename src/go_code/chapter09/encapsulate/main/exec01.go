package main

import (
	"fmt"
	"go_code/chapter09/encapsulate/model"
)

func main() {
	ac1 := model.NewAccount()
	//设置信息
	ac1.SetAccount("lemon616")
	ac1.SetPwd("888888")
	ac1.SetBalance(300)
	//获取账户
	fmt.Println("账户=", ac1.GetAccount())
	//获取余额
	fmt.Println("余额=", ac1.GetBalance("888888"))
}
