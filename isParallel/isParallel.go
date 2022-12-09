package main

import "fmt"

var(
	A1 = [2]int{}
	B1 = [2]int{}

	A2 = [2]int{}
	B2 = [2]int{}
)
func main(){
	twoPoint()

	var k1,k2 float64
	k1 = float64(( B1[1] - A1[1] ) /  ( B1[0] - A1[0] ))
	k2 = float64(( B2[1] - A2[1] ) /  ( B2[0] - A2[0] ))

	if k1 == k2 {
		fmt.Println("两线平行")
	}else {
		fmt.Println("两线不平行")
	}
}

func twoPoint(){
	fmt.Println("请输入线段L1的一点A1坐标")
	fmt.Scanln(&A1[0],&A1[1])
	fmt.Println("请输入线段L1的另一点B1坐标")
	fmt.Scanln(&B1[0],&B1[1])

	fmt.Println("请输入线段L2的一点A2坐标")
	fmt.Scanln(&A2[0],&A2[1])
	fmt.Println("请输入线段L2的另一点B2坐标")
	fmt.Scanln(&B2[0],&B2[1])
}