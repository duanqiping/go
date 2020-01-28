package main

import "fmt"

func main()  {
	str :="hello@iohioh"

	slice := str[6:]

	fmt.Println(slice)

	arr1 := []rune(str)

	arr1[0] = '北'

	fmt.Printf("arr1 = %c",arr1)
}
