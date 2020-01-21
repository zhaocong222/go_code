package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

//为了提高效率，通过我们方法和结构体的指针类型绑定
func (c *Circle) area2() float64 {
	//因为c是个指针，因此我们标准的访问其字段的方式是 (*c)
	//return 3.14 * (*c).radius * (*c).radius
	//底层自动优化，等价于下面的写法, (*c).radius等价c.radius
	c.radius = 1.0
	return 3.14 * c.radius * c.radius
}

func main() {
	// 1.声明一个结构体Circle，字段为radius
	// 2.声明一个方法area和Circle绑定，可以返回面积
	// 3.提示：画出area执行过程+说明
	circle := Circle{1.2}
	fmt.Println("面积是", circle.area())

	//创建一个circle变量
	var c Circle
	c.radius = 5.0
	//res2 := (&c).area2() //标准访问方式
	//编译器底层优化，也可以这么写
	res2 := c.area2()
	fmt.Println("c的面积是", res2)
	fmt.Println("radius是", c.radius)

}
