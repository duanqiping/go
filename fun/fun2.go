package main

import "fmt"

func fire()  {
	fmt.Println("fire")
}

//遍历切片的每个元素，通过给定函数进行元素访问
func visit(list []int, f func(int)){
	for _, v := range list{
		f(v)
	}
}

func main()  {
	var f func()
	f  = fire
	f()

	//将匿名函数保存到f1()中
	f1 := func(data int) {
		fmt.Println("hello",data)
	}
	//使用f1()调用
	f1(100)

	visit([]int{1,2,3,4}, func(v int) {
		fmt.Println(v)
	})
}
