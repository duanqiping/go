package main

import "fmt"

//编写一个函数调用九九乘法表
//1	1*1
//2 1*1 1*2 2*


func multi(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j<= i ; j++ {
			fmt.Printf("%v*%v = %v\t",j,i,i*j)
		}
		fmt.Println()
	}
}

func main()  {
	multi(8)
}
