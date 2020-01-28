package main

import "fmt"

func main()  {
	var arr = [...]int{10,20,30,40,50}

	slice := arr[1:3]
	len_slice := len(slice)
	for i := 0; i<len_slice; i++ {
		//fmt.Println(slice[i])
		fmt.Printf("slice[%v] = %v\n",i,slice[i])
	}
	fmt.Println()

	for i,v := range slice{
		fmt.Printf("slice[%v] = %v\n",i,v)
	}

	var  slice3  = []int{100,200,300}
	slice3 =append(slice3,400,500,600)
	fmt.Println("slice3 = ",slice3)

	slice3 = append(slice3,slice3...)
	fmt.Println("slice3 = ",slice3)

	var slice4 = []int{1,2,3,4,5}
	var slice5 = make([]int,10)
	copy(slice5,slice4)
	fmt.Println("slice5 = ",slice5)

}


