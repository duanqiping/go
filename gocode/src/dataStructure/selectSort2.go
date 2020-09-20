package main

import "fmt"

//从大到小排序
func selectSort(arr *[6]int)  {
	//(*arr)[1] = 600//标准写法
	//arr[1] = 601//优化后的写法

	for j := 0; j<len(arr); j++ {
		maxIndex := j
		maxValue := arr[j]

		for i := j+1;i < len(arr);i++ {
			if maxValue < arr[i] {
				maxValue = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j],arr[maxIndex] = arr[maxIndex],arr[j]
		}
		fmt.Println(arr)
		fmt.Println()
	}
}

func main()  {

	arr := [6]int{10,30,34,60,25,28}

	selectSort(&arr)

	fmt.Println(arr)
}
