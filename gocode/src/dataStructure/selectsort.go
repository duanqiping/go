package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SelectSort(arr *[80000]int)  {
	length := len(arr)
	for i := 0;i<length-1; i++ {
		maxValue := arr[i]
		maxIndex := i

		//2. 遍历后面 1---[len(arr) -1] 比较
		for j := i+1;j < length;j++ {
			if maxValue < arr[j] {
				maxValue = arr[j]
				maxIndex = j
			}
		}
		//交换
		if maxIndex != i {
			arr[i],arr[maxIndex] = maxValue,arr[i]
		}
	}
}

func main()  {
	//arr := [6]int{1,-23,44,22,2,12}
	//
	//SelectSort(&arr)
	//
	//fmt.Println(arr)

	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}

	//fmt.Println(arr)
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时=%d秒", end - start)
	fmt.Println("main函数")


}