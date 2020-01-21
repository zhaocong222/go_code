package main

import "fmt"

func main() {
	//求出每个班级平均分，以及所有班级平均分
	var sorce [3][5]float64

	for i, v := range sorce {
		for j, _ := range v {
			fmt.Printf("请输入%d班，第%d个学生的成绩", i+1, j+1)
			fmt.Scanln(&sorce[i][j])
		}
	}

	//求出平均分和所有班级平均分
	totalSum := 0.0
	person := 0
	for i, v := range sorce {
		sum := 0.0
		for _, val := range v {
			sum += val
			person++
		}
		totalSum += sum
		fmt.Printf("第%d班的总分为%v,平均分为%v\n", i, sum, sum/float64(len(v)))
	}
	fmt.Printf("所有班的总分为%v,所有班的平均分为%v\n",
		totalSum, totalSum/float64(person))

}
