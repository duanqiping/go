package main

import "fmt"

type A struct {
	Name string
	age int
}
type B struct {
	Name string
	Score float64
}

type C struct {
	A
	B
}

type D struct {
	a A
}

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}
type Monster struct  {
	Name string
	Age int
}

type E struct {
	Monster
	int
	n int
}

func main()  {
	var c C
	//如果c 没有Name字段，而A 和 B有Name, 这时就必须通过指定匿名结构体名字来区分
	//所以 c.Name 就会包编译错误， 这个规则对方法也是一样的！
	c.A.Name = "tom" // error
	fmt.Println("c",c)

	var d D
	d.a.Name = "jake"
	fmt.Println("d",d)

	tv := TV{Goods{"戴尔笔记本",49999},Brand{"联想笔记本","北京"}}
	fmt.Println("tv",tv)

	tv2 := TV2{&Goods{"戴尔笔记本1",49999},&Brand{"联想笔记本1","北京1"}}
	fmt.Println("tv2",tv2)
	fmt.Println("tv2",*tv2.Goods,*tv2.Brand)

	//演示一下匿名字段时基本数据类型的使用
	var e E
	e.Name = "狐狸精"
	e.Age = 300
	e.int = 20
	e.n = 40
	fmt.Println("e=", e)

}
