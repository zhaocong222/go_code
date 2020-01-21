package main

import "fmt"

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type D struct {
	Sex string
}

type C struct {
	A
	B
	d D //如果有名结构体，就为组合关系，访问时需要指明
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV3 struct {
	//用指针效率高，因为结构体是值传递
	*Goods
	*Brand
}

func main() {
	//如果嵌套的A,B两个匿名结构体，存在相同字段名，需要指明 (对方法也是一样)
	var c C
	//c.Name = "tom" //error
	c.A.Name = "hh"
	c.d.Sex = "男"

	fmt.Println(c.A.Name)
	fmt.Println(c.d.Sex)

	//声明时直接赋值
	tv := TV{
		Goods{"电视机", 5000},
		Brand{"海尔", "山东青岛"},
	}
	//也可以如下
	tv2 := TV{
		Goods{
			Name: "电视机",
			//不指定价格
		},
		Brand{
			Address: "山东青岛",
		},
	}

	//因为tv3结构体里面 2个匿名结构体指针，所以传入地址，&Goods....
	tv3 := TV3{
		&Goods{
			Price: 34534,
		},
		&Brand{
			Address: "北京",
		},
	}

	fmt.Println(tv)
	fmt.Println(tv2)

	//等价 fmt.Println(*(tv3.Goods), *(tv3.Brand))
	fmt.Println(*tv3.Goods, *tv3.Brand)

}
