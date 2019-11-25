package main

import "fmt"

func main()  {
	n := add(4,5)
	m := sub(2,3)
	p := zero(2,3)

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(p)
	a,b := typeTwoValues()
	fmt.Println(a,b)

	fmt.Printf("%T\n", add) // "func(int, int) int"
	fmt.Printf("%T\n", sub) // "func(int, int) int"
	fmt.Printf("%T\n", zero) // "func(int, int) int"
}
func zero(int, int) int { return 0 }

func add(x int,y int) int {
	return x+y
}

func sub(x,y int) (z int)  {
	z = x - y;return
}

func typeTwoValues() (int,int) {
	return 1,2
}
