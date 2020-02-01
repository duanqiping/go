package main

import "fmt"

func main()  {
	a := [3]int{1,2,3}

	b := a[1:2]

	c := append(b,8,9,10,11)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
