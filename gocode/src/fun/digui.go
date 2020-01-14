package main

import "fmt"

//题2：求函数值已知 f(1)=3; f(n) = 2*f(n-1)+1; 请使用递归的思想编程，求出 f(n)的值?

func f(n float64) float64  {
	if n == 1{
		return 3
	}else{
		return 2*f(n-1)+1
	}
}

func main()  {
	//println("f(n) =",f(1))
	//println("f(n) =",f(2))
	//println("f(n) =",f(3))
	//println("f(n) =",f(100))

	fmt.Printf("f(n) = %v\n",f(3))
	fmt.Printf("f(n) = %v\n",f(50))
}