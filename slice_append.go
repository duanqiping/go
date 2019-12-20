package main

// 导入系统包
import (
	"fmt"
)

func main() {

	var a = []int{1,2,3}
	a = append([]int{0}, a...) // 在开头添加1个元素
	a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片

	fmt.Println(a)
	
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	// copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置

	fmt.Println(slice1)
	fmt.Println(slice2)
}