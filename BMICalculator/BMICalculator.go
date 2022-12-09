package main

import "fmt"

var (

	name string
	sex string
	sexRate float64
	height float64
	weight float64
	age int

	bmi float64
	fatRate float64
	averageBMI float64 = 0.0
	averageFatRate float64 = 0.0
	SumFatRate float64 =0.0
	peopleCount int = 0
	isContinue bool

)//用结构体声明变量更好一些吗


func main(){
	fmt.Println("BMICalculator is running....")

	for isContinue = true; isContinue; {
		getPerson()
		peopleCount ++

		//bmi和体脂率的计算公式
		bmi = weight/(height * height)
		fatRate = (1.2 * bmi + 0.23 * float64(age) - 5.4 - 10.8 * sexRate) / 100

		//求平均体脂率
		SumFatRate += fatRate
		averageFatRate = SumFatRate / float64(peopleCount)
		calculate()
		fmt.Println()
		fmt.Printf("总人数是%d，平均体脂率为%2.3f\n",peopleCount,averageFatRate)
		fmt.Println("*************************************")

		exitBMICalculator()
	}

}


func getPerson(){
	fmt.Printf("姓名：")
	fmt.Scanln(&name)
	fmt.Printf("性别：")
	fmt.Scanln(&sex)
	fmt.Printf("身高：")
	fmt.Scanln(&height)
	fmt.Printf("体重：")
	fmt.Scanln(&weight)
	fmt.Printf("年龄：")
	fmt.Scanln(&age)
	if sex == "male" {
		sexRate = 1
	}else {
		sexRate = 0
	}
}


func calculate(){
	fmt.Println("*************************************")
	//业务逻辑就简单写了
	if sex == "male" {
		if fatRate >= 0.15 && fatRate <= 0.18{
			fmt.Printf("BMI:%2.3f,体脂率：%2.3f\n 恭喜在正常范围内",bmi,fatRate)
		}else {
			fmt.Printf("BMI:%2.3f,体脂率：%2.3f\n 不正常范围内，需要留意了哦",bmi,fatRate)
		}
	}else {
		if fatRate >= 0.2 && fatRate <= 0.25{
			fmt.Printf("BMI:%2.3f,体脂率：%2.3f\n 恭喜在正常范围内",bmi,fatRate)
		}else {
			fmt.Printf("BMI:%2.3f,体脂率：%2.3f\n 不正常范围内，需要留意了哦",bmi,fatRate)
		}
	}
}

func exitBMICalculator(){
	fmt.Println("按 Y 继续，按 Q 退出")
	var temp string
	fmt.Scanln(&temp)
	if temp == "Y" || temp == "y" {
		isContinue = true
	}else if temp == "Q" || temp == "q"{
		isContinue = false
	}
}