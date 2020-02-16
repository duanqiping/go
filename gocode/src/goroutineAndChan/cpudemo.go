package main

import (
	"fmt"
	"runtime"
)

func main()  {
	cpuNum := runtime.NumCPU()

	fmt.Println(cpuNum)//8

	runtime.GOMAXPROCS(cpuNum - 4)

	cpuNum2 := runtime.NumCPU()

	fmt.Println(cpuNum2)//8


}
