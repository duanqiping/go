package main

import "fmt"

type Person3 struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int//指针
	slice []int //切片
	map1 map[string]string //map 隐射
}

type Monster1 struct {
	Name string
	Age int
}

func main()  {
	//定义结构体变量
	var p3 Person3
	fmt.Println(p3)

	if p3.ptr == nil {
		fmt.Println("ok1")
	}
	if p3.slice == nil {
		fmt.Println("ok2")
	}

	if p3.map1 == nil {
		fmt.Println("ok3")
	}

	p3.slice = make([]int,10)
	p3.slice[0] = 100
	fmt.Println(p3)

	//使用map, 一定要先make
	p3.map1 = make(map[string]string)
	p3.map1["key1"] = "tom~"
	fmt.Println(p3)

	//不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改，
	//不影响另外一个, 结构体是值类型
	var monster1 Monster1
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1 //结构体是值类型，默认为值拷贝
	monster2.Name = "青牛精"

	fmt.Println("monster1=", monster1) //monster1= {牛魔王 500}
	fmt.Println("monster2=", monster2) //monster2= {青牛精 500}
}
