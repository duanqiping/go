package main

import "fmt"

func main()  {

	var array01 = [3]int{1,2,3}
	fmt.Println(array01)

	var array02 = [...]int{4,5,6}
	fmt.Println(array02)

	var array03 = [...]int{1:4,2:5,3:6}
	fmt.Println(array03)

	//类型推导
	strArr04 := [...]string{1: "tom", 0: "jack", 2:"mary"}
	fmt.Println("strArr05=", strArr04)
}
