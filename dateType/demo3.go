package main

import "fmt"

func test() bool {
	fmt.Println("ok")
	return true
}

func main()  {

	var n1 int = 10
	if n1 < 0 && test() {
		fmt.Println("ok1")
	}

	if n1 > 0 || test() {
		fmt.Println("ok1")
	}
}
