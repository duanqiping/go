package main

import "fmt"

func main()  {
	var hero map[string]string

	hero = make(map[string]string,1)

	hero["num1"] = "宋江"
	hero["num2"] = "武松"
	hero["num3"] = "鲁智深"
	hero["num0"] = "鲁智深2"

	fmt.Println(hero)
}
