package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	score int
}

//小学生
type Pupil struct {
	Student //嵌入Student匿名结构体,有了其属性
}

//大学生
type Graduate struct {
	Student //嵌入Student匿名结构体,有了其属性
}

//小学生和大学生共有的方法也绑定到*student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v,年龄=%v,成绩=%v\n", stu.Name, stu.Age, stu.score)
}

//设置分数
func (stu *Student) SetScore(score int) {
	stu.score = score
}

//给 *Student增加一个方法，那么Pupil和Graduate都可以使用该方法
func (stu *Student) getSum(n1 int, n2 int) {
	fmt.Println(n1 + n2)
}

//小学生考试-特有方法
func (p *Pupil) testing() {
	fmt.Println("小学生考3门")
}

//大学生考试-特有方法
func (g *Graduate) testing() {
	fmt.Println("大学生考5门")
}

func main() {
	//值类型
	// var pupil = Pupil{}

	//也可以写成如下,下面就是指针类型
	var pupil = &Pupil{}
	pupil.Student.Name = "tome"
	pupil.Student.Age = 21
	pupil.testing() //考试
	pupil.SetScore(80)
	pupil.Student.ShowInfo()
}
