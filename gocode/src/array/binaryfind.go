package main

import (
	"fmt"
)

//二分查找
//1数字是有序的，2从中间的数开始找 MiddleVal MiddleIndex
func binaryFind(arr [8]int,leftIndex int,rightIndex int, findVal int)  {

	if leftIndex >= rightIndex {
		fmt.Println("找不到")
		return
	}

	//先找到中间下标
	middleIndex := (leftIndex + rightIndex)/2

	if findVal == arr[middleIndex] {
		fmt.Println("找到了，数组下标是",middleIndex)
	}
	if findVal > arr[middleIndex] {
		//再中间数的右边找
		binaryFind(arr,middleIndex+1,rightIndex,findVal)
	}
	if findVal < arr[middleIndex] {
		//从中间的数左边找
		binaryFind(arr,leftIndex,middleIndex-1,6)
	}

	//bLen := len(arr)
	//MiddleVal := arr[math.Ceil(int(bLen/2]))]

}

func main()  {
	arr := [...]int{1,2,3,10,23,300,400,500}

	binaryFind(arr,0,len(arr)-1,400)

}