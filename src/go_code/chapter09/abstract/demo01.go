package main

import "fmt"

type Account struct {
	Account string
	Pwd     string
	Balance float64
}

//方法
//1.存款
func (account *Account) Deposite(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money

	fmt.Println("存款成功")
}

//2.取款
func (account *Account) WithDraw(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//取款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance -= money

	fmt.Println("取款成功")
}

//3.查询余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为%v,余额为%v\n", account.Account, account.Balance)
}

func main() {
	//测试一把
	ac1 := Account{
		Account: "工商1111",
		Pwd:     "666666",
		Balance: 100.0,
	}

	ac1.Query("666666")
	ac1.Deposite(200.0, "666666")
	ac1.Query("666666")

}
