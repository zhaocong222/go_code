package service

import (
	"go_code/customerManage/model"
)

//增删改查
type CustomerService struct {
	customer []model.Customer //切片
	//声明一个字段，表示当前切片含有多少个客户,还可以作为新客户的id+1
	custimerNum int
}

func (this *CustomerService) List() []model.Customer {
	return this.customer
}

//添加客户到customer切片
func (this *CustomerService) Add(customer model.Customer) bool {

	//分配id，用添加的顺序
	this.custimerNum++
	customer.Id = this.custimerNum
	this.customer = append(this.customer, customer)
	return true
}

//根据id删除客户(从切片中)
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}

	//找到删除
	this.customer = append(this.customer[:index], this.customer[index+1:]...)

	return true
}

//根据id查找客户在切片中对应下标没如果没有返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历切片this.customer 切片
	for i := 0; i < len(this.customer); i++ {
		if id == this.customer[i].Id {
			//找到
			index = i
			break
		}
	}

	return index
}

func NewCustomerService() *CustomerService {
	//初始化客户
	customerService := &CustomerService{}
	customerService.custimerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "za@qq.com")
	customerService.customer = append(customerService.customer, customer)

	return customerService
}
