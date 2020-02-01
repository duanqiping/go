package main

import "fmt"

type Point struct {
	x int
	y int
}
type Point2 struct {
	x int
	y int
}

func main()  {
	var a interface{}
	var point = Point{1,2}
	a = point

	var b Point
	b = a.(Point)

	fmt.Println(a)
	fmt.Println(b)

	//类型断言的其它案例
	//var x interface{}
	//var z float32 = 1.1
	//x = z
	//
	//var y float32
	//y = x.(float32)
	//fmt.Printf("y 的类型是 %T 值是=%v", y, y)

	//类型断言(带检测的)
	var x interface{}
	var b2 float32 = 2.1
	x = b2  //空接口，可以接收任意类型
	// x=>float32 [使用类型断言]

	//类型断言(带检测的)
	if y, ok := x.(float32); ok {
		fmt.Printf("ok 的类型是 %T 值是=%v", ok, ok)
		fmt.Println("convert success")
		fmt.Printf("y 的类型是 %T 值是=%v", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")


}
