package main

import (
	"fmt"       // 导入 fmt 包，打印字符串是需要用到
)


func main() {
	var a  = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])

	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])

	//[11 12 13 14 15]
	//[21 22 23 24 25 26 27 28 29 30]
	//[1 2]
}