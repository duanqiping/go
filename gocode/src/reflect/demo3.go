package main

import (
	"fmt"
	"reflect"
)

//通过反射，修改,
// num int 的值
// 修改 student的值

func reflect03(b interface{}) {
	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	// 看看 rVal的Kind是
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	//3. rVal
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
	rVal.Elem().SetInt(20)
}


func main()  {
	var num  =  10
	reflect03(&num)
	fmt.Println("num=", num) // 20
}
