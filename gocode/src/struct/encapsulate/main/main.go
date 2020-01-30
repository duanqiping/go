package main

import (
	"fmt"
	"struct/encapsulate/model"
)

func main()  {
	person := model.NewPerson("qiping")

	person.SetAge(29)
	person.SetSal(30000)

	fmt.Println(person)
}
