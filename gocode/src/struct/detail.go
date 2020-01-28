package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp,rightUp Point
}

type Rect2 struct {
	leftUp,rightUp *Point
}

func main()  {
	r1 := Rect{Point{3,4},Point{5,6}}

	//r1有4个int, 在内存中连续分布
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightUp.x 地址=%p r1.rightUp.y 地址=%p \n",
		&r1.leftUp.x,&r1.leftUp.y,&r1.rightUp.x,&r1.rightUp.y)

	r2 := Rect2{&Point{7,8},&Point{10,11}}
	fmt.Printf("r2.leftUp.x 地址=%p r2.leftUp.y 地址=%p r2.rightUp.x 地址=%p r2.rightUp.y 地址=%p \n",
		&r2.leftUp.x,&r2.leftUp.y,&r2.rightUp.x,&r2.rightUp.y)

	//他们指向的地址不一定是连续...， 这个要看系统在运行时是如何分配
	fmt.Printf("r2.leftUp 地址=%p r2.rightUp 地址=%p\n",
		r2.leftUp,r2.rightUp)
}
