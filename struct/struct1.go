package main

import . "fmt"

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

type People struct {
	name  string
	child *People
}

func main()  {

	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}

	Println(addr)
	Println(relation)
	Println(relation.name)
	Println(relation.child)
}
