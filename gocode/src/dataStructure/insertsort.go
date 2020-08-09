package main

import "fmt"

func InsertSort(arr *[6]int) {
	//假设第一个数是有序的

	//完成第一次，给第二个元素找到合适的位置并插入
	for i := 1; i<len(arr); i++ {

		//insertValue := arr[i]
		//insertIndex := i-1

		insertVal := arr[i]
		insertIndex := i - 1 // 下标

		//从大到小  这个方法失败的 原因是 if ，应该改成 for
		//if insertIndex >= 0 && arr[insertIndex] < insertVal {
		//	arr[insertIndex + 1] = arr[insertIndex]
		//	insertIndex--
		//}

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
			insertIndex--
		}

		//插入
		if insertIndex  != i {
			arr[insertIndex+1] = insertVal
		}

	}

}

//从大到小排序
func main()  {

	arr := [6]int{1,-23,44,22,2,12}

	InsertSort(&arr)
	fmt.Println(arr)
}

//package main
//
//import (
//	"fmt"
//)

//func InsertSort(arr *[6]int) {
//
//	//完成第一次，给第二个元素找到合适的位置并插入
//
//	for i := 1; i < len(arr); i++ {
//
//		insertVal := arr[i]
//		insertIndex := i - 1 // 下标
//
//		//从大到小
//		for insertIndex >= 0 && arr[insertIndex] < insertVal {
//			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
//			insertIndex--
//		}
//		//插入
//		if insertIndex + 1 != i {
//			arr[insertIndex + 1] = insertVal
//		}
//		//fmt.Printf("第%d次插入后 %v\n",i, *arr)
//	}
//}


//func main() {
//
//
//
//	arr := [6]int{0, 12, 56,  34, -1, 55}
//	InsertSort(&arr)
//	fmt.Println(arr)

	//var arr [80000]int
	//for i := 0; i < 80000; i++ {
	//	arr[i] = rand.Intn(900000)
	//}
	//
	////fmt.Println(arr)
	//start := time.Now().Unix()
	////fmt.Println("原始数组=", arr)
	//InsertSort(&arr)
	//end := time.Now().Unix()
	//
	//fmt.Println("main 函数")
	//fmt.Printf("插入排序耗时%d秒", end-start)
	////fmt.Println(arr)
//}