package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now()
	for i := 1; i <= 40; i++ {
		result = fibonacci1(i)
		fmt.Printf("数列第 %d 位: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("程序的执行时间为: %s\n", delta)
}
func fibonacci1(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci1(n-1) + fibonacci1(n-2)
	}
	return
}