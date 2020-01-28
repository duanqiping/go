package main

import "fmt"

type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
}

func main() {
	// 张老太养了20只猫猫:一只名字叫小白,今年3岁,白色。还有一只叫小花,
	// 今年100岁,花色。请编写一个程序，当用户输入小猫的名字时，就显示该猫的名字，
	// 年龄，颜色。如果用户输入的小猫名错误，则显示 张老太没有这只猫猫。

	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"

	var cat2 Cat
	cat2.Name = "小花"
	cat2.Age = 100
	cat2.Color = "花色"

	fmt.Println(cat1)
	fmt.Println(cat2)
	fmt.Println(cat2.Name)

}