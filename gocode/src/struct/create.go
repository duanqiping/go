package main

import "fmt"

type Person struct {
	Age int
	Name string
}

func main()  {

	//方式1
	var p1 Person
	p1.Age = 10
	p1.Name = "tom"
	fmt.Println(p1)

	//方式2
	p2 := Person{20,"jack"}
	fmt.Println(p2)

	//方式3
	//var p3 *Person = new(Person)
	var p3 = new(Person)
	//(*p3).Name = "tom^"
	p3.Name = "tom^^"
	//(*p3).Age = 30
	p3.Age = 40
	fmt.Println(p3)

	//方式4
	var p4 = &Person{}
	p4.Name = "jack^^"
	p4.Age = 50
	fmt.Println(p4)

}
