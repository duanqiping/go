package main

import "fmt"

func main()  {
	a := 100

	fmt.Println("a的地址是：",&a)

	var ptr *int = &a
	fmt.Println("ptr 指向的值是：",ptr)

	*ptr = 200

	fmt.Println("a 的值是：",a)

	var n int
	var i int = 10
	var j int = 12

	//n = i > j ? j : i  //go不支持三元运算

	fmt.Println("n 的值是：",n)

}