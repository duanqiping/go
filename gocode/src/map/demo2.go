package main

import "fmt"

//定义一个学生结构体
type Stu struct {
	Name string
	age int
	address string
}


func modiy(map1 map[int]int)  {
	map1[3] = 20
}

func main()  {
	//map是引用类型，遵守引用类型传递的机制，在一个函数接收map，
	//修改后，会直接修改原来的map

	map1 := make(map[int]int,2)

	map1[1] = 100
	map1[2] = 101
	map1[4] = 102

	modiy(map1)
	fmt.Println(map1)

	//map的value 也经常使用struct 类型，
	//更适合管理复杂的数据(比前面value是一个map更好)，
	//比如value为 Student结构体 【案例演示，因为还没有学结构体，体验一下即可】
	//1.map 的 key 为 学生的学号，是唯一的
	//2.map 的 value为结构体，包含学生的 名字，年龄, 地址

	map2 := make(map[int64]Stu)
	stu1 := Stu{"tom",15,"江西"}
	stu2 := Stu{"tom1",16,"武汗"}
	stu3 := Stu{"tom2",17,"上海"}

	map2['1'] = stu1
	map2['2'] = stu2
	map2['3'] = stu3

	fmt.Println(map2)

	//遍历各个学生信息
	for k, v := range map2 {
		fmt.Printf("学生的编号是%v \n", k)
		fmt.Printf("学生的名字是%v \n", v.Name)
		fmt.Printf("学生的年龄是%v \n", v.age)
		fmt.Printf("学生的地址是%v \n", v.address)
		fmt.Println()
	}
}
