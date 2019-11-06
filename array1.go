package main

import (
	"fmt"       // 导入 fmt 包，打印字符串是需要用到
)


func main() {
	var a [3]int             // 定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"

	//team := [...]string{hammer,soldier,mum}//.\array1.go:29:22: undefined: hammer
	for k, v := range team {
		fmt.Println(k, v)
	}
}