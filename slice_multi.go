package main

import "fmt"

func main() {
	// 创建一个整型切片
	slice := [][]int{{10}, {100, 200}}
	// 为第一个切片追加值为 20 的元素
	slice[0] = append(slice[0], 20)

	fmt.Println(slice)

	//[11 12 13 14 15]
	//[21 22 23 24 25 26 27 28 29 30]
	//[1 2]
}