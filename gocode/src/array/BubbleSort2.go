package main

import "fmt"

//冒泡排序 从小到大排序
func BubbleSort2(arr *[6]int)  {

	for i:=0;i<len(*arr)-1 ;i++  {
		for j:=0; j<len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
		fmt.Println(*arr)
	}
}

func main()  {
	var arr = [6]int{1,2,-10,4,-3,0}

	fmt.Println(arr)
	BubbleSort2(&arr)
	fmt.Println(arr)
}
