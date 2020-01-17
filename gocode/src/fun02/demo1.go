package main

import (
	"fmt"
	"fun02/utils"
)

//全局变量
var age = test()

func test()  int{
	fmt.Println("test 函数")
	return 10
}

//init函数,通常可以在init函数中完成初始化工作
func init()  {
	fmt.Println("init 函数")
}

func main()  {
	fmt.Println("main age=",age)
	fmt.Println("main 函数")

	utils.Test01()
}
