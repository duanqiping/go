package main

import "fmt"

func main()  {
	//使用常规的for循环遍历切片
	var sliceArray = [5]int{1,2,3,100,200}

	fmt.Println(sliceArray)

	arrayLen := len(sliceArray)
	for i := 0; i<arrayLen;i++  {
		fmt.Printf("sliceArray[%d]=%d\n",i,sliceArray[i])
	}

	slice := sliceArray[1:]
	fmt.Println(slice)
	for i,v := range slice {
		fmt.Printf("i=%d,v=%d\n",i,v)
	}
	slice2 := slice[1:2]
	fmt.Println(slice2)
	slice2[0] = 500
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(sliceArray)

	//append的用法
	slice3 := append(slice2,10,20)
	fmt.Println(slice3)

	slice3 = append(slice3,slice3...)
	fmt.Println(slice3)

	//copy的用法
	var slice4 = []int{1,2,3,4}
	var slice5 = make([]int,5)
	copy(slice5,slice4)
	fmt.Println(slice4)
	fmt.Println(slice5)

}