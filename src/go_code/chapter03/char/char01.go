package main
import (
	"fmt"
)

func main()  {
	var c1 byte = 'a'
	var c2 byte = '0' //字符0
	
	//当我们直接输出byte值，就是输出了对应的字符的码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//如果我们希望输出对应的字符，需要格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	//var c3 byte = '北' // constant 21271 overflows byte溢出
	var c3 int = '北'
	fmt.Printf("c3=%c\n", c3)

	//字符类型是可以进行运算的，相当于一个整数，运算时是按照码值运行
	var n1 = 10 + 'a' // 10 +97 => 107
	fmt.Println("n1=", n1)
}