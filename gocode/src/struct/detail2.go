package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}
type B struct {
	Num int
}

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main()  {
	var a A
	var b B
	a = A(b)
	fmt.Println(a,b)

	//1.创建一个monster 变量
	monster := Monster{"牛魔王",500,"牛魔拳"}
	jsonStr,err := json.Marshal(monster)
	if err == nil {
		fmt.Printf("%v",string(jsonStr))
	}else {
		fmt.Println("json转化 失败...")
	}


}
