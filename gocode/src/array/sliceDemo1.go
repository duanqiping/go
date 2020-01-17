package main

import "fmt"

func main()  {
	//演示切片的基本使用
	var intArr  = [...]int{1, 22, 33, 66, 99}

	slice := intArr[1:3]

	fmt.Println("slice的 值",slice)

	//元素个数
	fmt.Println(len(slice))

	//元素容量
	fmt.Println(cap(slice))

	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址=%p slice[0==%v\n", &slice[0], slice[0])
	slice[1] = 34
	fmt.Println()
	fmt.Println()
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice) //  22, 33
}