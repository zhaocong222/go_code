package model

import "fmt"

type account struct {
	account string
	balance float64
	pwd     string
}

func NewAccount() *account {
	return &account{}
}

func (a *account) SetAccount(aname string) {
	if len(aname) >= 6 && len(aname) <= 10 {
		a.account = aname
	} else {
		fmt.Println("输入的账号必须在6-10位之间")
	}
}

func (a *account) SetBalance(balance float64) {
	if balance >= 20 {
		a.balance = balance
	} else {
		fmt.Println("余额不得小于20")
	}
}

func (a *account) SetPwd(pwd string) {
	if len(pwd) == 6 {
		a.pwd = pwd
	} else {
		fmt.Println("密码必须是6位")
	}
}

func (a *account) GetAccount() string {
	return a.account
}

func (a *account) GetBalance(pwd string) float64 {
	if pwd == a.pwd {
		return a.balance
	} else {
		fmt.Println("密码不正确")
		return 0
	}
}

func (a *account) GetPwd() string {
	return a.pwd
}
