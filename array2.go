package main

import (
	"fmt"       // 导入 fmt 包，打印字符串是需要用到
)


func main() {
	// 声明两个二维整型数组
	var array1 [2][2]int
	var array2 [2][2]int
	// 为array2的每个元素赋值
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	// 将 array2 的值复制给 array1
	array1 = array2

	// 将 array1 的索引为 1 的维度复制到一个同类型的新数组里
	var array3 [2]int = array1[1]
	// 将数组中指定的整型值复制到新的整型变量里
	var value int = array1[1][0]

	fmt.Println(array3)
	fmt.Println(value)

	//[30 40]
	//30
}