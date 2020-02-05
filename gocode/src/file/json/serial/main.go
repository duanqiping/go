package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday string `json:"birthday"`
	Sal float64 `json:"sal"`
	Skill string `json:"skill"`
}

func testStruct()  {
	monster := Monster{
		Name:"牛魔王",
		Age:500,
		Birthday:"8月13",
		Sal:5000,
		Skill:"牛魔拳",
	}
	data,err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("结构体序列化失败",err)
	}
	//输出序列化后的结果
	fmt.Printf("结构体序列化后的结果：%v \n",string(data))
}

func testMap()  {
	var a map[string]interface{}
	a  = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 10
	a["address"] = "洪崖洞"

	data,err := json.Marshal(a)
	if err != nil {
		fmt.Println("结构体序列化失败",err)
	}
	//输出序列化后的结果
	fmt.Printf("结构体序列化后的结果：%v \n",string(data))
}

func testSlice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用map前，需要先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	//使用map前，需要先make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥","夏威夷"}
	slice = append(slice, m2)

	data,err := json.Marshal(slice)
	if err != nil {
		fmt.Println("结构体序列化失败",err)
	}
	//输出序列化后的结果
	fmt.Printf("结构体序列化后的结果：%v \n",string(data))
}

//对基本数据类型序列化，对基本数据类型进行序列化意义不大
func testFloat64() {
	var num1 float64 = 2345.67

	//对num1进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("num1 序列化后=%v\n", string(data))
}

func main()  {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
