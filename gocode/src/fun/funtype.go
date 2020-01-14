package main

import "fmt"

//在Go中，函数也是一种数据类型，
//可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用

func getSum(n1,n2 int) int {
	return n1 + n2
}

type myFunvar func(int,int) int
func myFun2(funval myFunvar,n1 int,n2 int) int {
	return funval(n1,n2)
}

func getSubSum(n1,n2 int) (sum int,sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//可变参数的使用
func changeParams(n1 int,argus... int) int {
	sum := 0
	sum += n1
	for i := 0; i<len(argus);i++ {
		sum += argus[i]
	}
	return sum
}


func main()  {
	sum := getSum

	fmt.Printf("res的类型是%T sum的类型是%T\n",sum,getSum)

	res := sum(1,2)
	fmt.Println(res)

	sum2 := myFun2(getSum,40,60)
	fmt.Printf("sum2的值是%v\n",sum2)

	type myint int

	var num1 myint
	var num2 int

	num2 = 20
	num1 = myint(num2)

	fmt.Println("num1=", num1, "num2=",num2)

	sumTotal,sub := getSubSum(10,20)
	fmt.Printf("sum的值是%v sub的值是%v\n",sumTotal,sub)

	sumTotal2 := changeParams(1,2,3,4,-5)
	fmt.Printf("sumTotal2的值是%v\n",sumTotal2)

}