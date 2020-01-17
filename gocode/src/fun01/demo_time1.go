package main

import (
	"fmt"
	"time"
)

//100000  打印循环的耗时
var temp int64 = 1

func a()  {
	var i int64 = 1
	for ; i<1000000 ; i++ {
		if i < 10 {
			temp *= i
		}

	}
}

func main()  {
	st := time.Now()
	a()
	et := time.Since(st)//等价于 T := time.Now().Sub(startT)
	fmt.Println("a 函数用时 ",et,temp)

	//st := time.Now().Unix()
	//a()
	//et := time.Now().Unix()

	//fmt.Println("a 函数用时 ",et - st)
}
