package main

import "fmt"

func  main()  {

	//去掉小数保留整数部分
	fmt.Println(10/4)

	var n1 float32 = 10/4
	fmt.Println(n1)

	var n2 float32 = 10.0/4
	fmt.Println(n2)

	var i int8 = 10
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)

	if i>0{
		fmt.Println("ok")
	}
}