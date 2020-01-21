package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func main() {
	var p1 Person
	p1.Age = 10
	p1.Name = "小明"
	var p2 *Person = &p1

	fmt.Println((*p2).Age) // 10
	fmt.Println(p2.Age)    // 10
	p2.Name = "tom~"
	fmt.Println("p2.Name=%v,p1.Name=%v\n", p2.Name, p1.Name)    //tom~, tom~
	fmt.Println("p2.Name=%v,p1.Name=%v\n", (*p2).Name, p1.Name) //tom~,tom~

	//输出地址
	fmt.Printf("p1的地址%p\n", &p1)
	fmt.Printf("p2的地址%p p2的值%p\n", &p1, p2)

}
