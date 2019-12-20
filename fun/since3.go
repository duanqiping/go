package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 1; i < LIM; i++ {
		result = fibonacci2(i)
		fmt.Printf("数列第 %d 位: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("程序的执行时间为: %s\n", delta)
}
func fibonacci2(n int) (res uint64) {
	// 记忆化：检查数组中是否已知斐波那契（n）
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs[n] = res
	return
}