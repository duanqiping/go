package main

import "fmt"
var fun1 = func(num1,num2 int) int {
	return num1*num2
}

//fun1就是一个全局匿名函数

func main()  {

	//案例演示,求两个数的和， 使用匿名函数的方式完成
	res := func(num1,num2 int) int {
		return num1 + num2
	}(10,20)

	fmt.Println(res)

	//将匿名函数func (n1 int, n2 int) int赋给 a变量
	//则a 的数据类型就是函数类型 ，此时,我们可以通过a完成调用
	a := func(num1,num2 int) int {
		return num1 + num2
	}
	res2 := a(30,40)
	fmt.Println(res2)

	//全局匿名函数的使用
	res4 := fun1(4,5)
	fmt.Println(res4)
}


