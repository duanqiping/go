package main

import (
	"fmt"
	"struct/encapexercise/model"
)

func main()  {
	account := model.NewAccount("11111111","222222",66.66)

	if account != nil {
		fmt.Println("创建成功...")
		fmt.Println(account)
	}else{
		fmt.Println("创建失败...")
	}
}
