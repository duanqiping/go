package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var count byte = 0

	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n =",n)
		if(n == 99){
			break
		}
		count++
	}
	fmt.Println("生成 99 一共循环了",count)

	//100以内的数求和，求出 当和 第一次大于20的当前数
	sum := 0
	for i := 1;i<=100;i++{
		sum += i
		if sum>20 {
			println("i的值为",i)
			break
		}
	}

	//实现登录验证，有三次机会，如果用户名为”张无忌” ,密码”888”提示登录成功，
	//否则提示还有几次机会.

	var name string
	var password string
	var len byte = 4
	var i byte  = 1

	for ; i<len;  {
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)

		fmt.Println("请输入密码")
		fmt.Scanln(&password)

		if name == "张无忌" && password == "888" {
			fmt.Println("登录成功")
			break
		} else {
			i++
			fmt.Printf("还有 %d 次机会",len-i)
		}
	}


}
