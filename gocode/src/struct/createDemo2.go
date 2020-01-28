package main

import "fmt"

type Person1 struct {
	Name string
	Age int
}


func main()  {
	var p1 Person1
	p1.Name = "小明"
	p1.Age = 12

	p2 := &p1
	p2.Name = "小明^^"
	fmt.Println(p1)
	fmt.Println(p2)
}
