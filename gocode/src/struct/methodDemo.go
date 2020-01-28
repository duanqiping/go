package main

import "fmt"

type Person4 struct {
	Name string
}

func (p4 Person4) jiSun(n int)  {
	sum := 0
	for i := 0; i < n;i++  {
		sum += i
	}
	fmt.Println("sum的值为：",sum)
}

func (p4 Person4) Speak()  {
	fmt.Println(p4.Name,"is a good man")
}

//给Person结构体添加getSum方法,可以计算两个数的和，并返回结果
func (p4 Person4) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var p4 Person4
	p4.Name = "Jake"
	p4.Speak()
	p4.jiSun(100)
	sum := p4.getSum(10,30)
	fmt.Println(sum)
}
