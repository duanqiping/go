package main

import "fmt"

//冒泡排序 从大到小排序
func BubbleSort(arr *[6]int)  {
	for j:=0; j<len(*arr)-1;j++ {
		arriLen := len(*arr)
		for i:=0; i<arriLen-1-j;i++  {
			if (*arr)[i] < (*arr)[i+1] {
				temp := (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}
		}
		fmt.Println(*arr)
	}
}

func main()  {
	var arr = [6]int{1,2,-10,4,-3,0}

	fmt.Println(arr)
	BubbleSort(&arr)
	fmt.Println(arr)
}
