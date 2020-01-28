package main

import "fmt"

type MethodUtils struct {
	//字段...
}

func (methodutils MethodUtils) PrintJuXing()  {
	for i := 0; i < 10; i++ {
		for j := 0;j < 8 ;j++  {
			fmt.Printf("%c",'*')
		}
		fmt.Println()
	}
}

//2)编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
func (mu MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
编写一个方法算该矩形的面积(可以接收长len，和宽width)，
将其作为方法返回值。在main方法中调用该方法，接收返回的面积值并打印
*/

func (mu MethodUtils) area(len float64, width float64) (float64) {
	return len * width
}

/*
编写方法：判断一个数是奇数还是偶数

*/

func (mu *MethodUtils) JudgeNum(num int)  {
	if num % 2 == 0 {
		fmt.Println(num, "是偶数..")
	} else {
		fmt.Println(num, "是奇数..")
	}
}

type Calcuator struct{
	Num1 float64
	Num2 float64
}
func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("运算符输入有误...")

	}
	return res
}

func main()  {
	/*
		1)编写结构体(MethodUtils)，编程一个方法，方法不需要参数，
		在方法中打印一个10*8 的矩形，在main方法中调用该方法。
	*/
	var methodutils MethodUtils
	methodutils.PrintJuXing()
	methodutils.Print2(3,20)

	methodutils.JudgeNum(11)

	//测试一下:
	var calcuator Calcuator
	calcuator.Num1 = 1.2
	calcuator.Num2 = 2.2
	res := calcuator.getRes('*')
	fmt.Println("res=", res)

}
