package main

import "fmt"

func main()  {
	//思路
	//1. 声明一个数组 var intArr[5] = [...]int {1, -1, 9, 90, 11}
	//2. 假定第一个元素就是最大值，下标就0
	//3. 然后从第二个元素开始循环比较，如果发现有更大，则交换

	var intArr = [...]int {1, -1, 9, 90, 11}

	var maxValue  = intArr[0]
	var MaxIndex  = 0

	arrayLen := len(intArr)
	for i := 1; i<arrayLen; i++ {
		if maxValue < intArr[i] {
			maxValue = intArr[i]
			MaxIndex = i
		}
	}

	fmt.Println("maxValue = ",maxValue)
	fmt.Println("MaxIndex = ",MaxIndex)
}
