package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("命令行的参数有", len(os.Args))
	//遍历
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}

}
