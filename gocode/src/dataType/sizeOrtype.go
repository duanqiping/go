package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var n int32 = 8

	fmt.Printf("字节大小：%d，数据类型：%T",unsafe.Sizeof(n),n)
}
