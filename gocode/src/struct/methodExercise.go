package main

import "fmt"

type Circle struct {
	dir float64
}

//2)声明一个方法area和Circle绑定，可以返回面积。
func (circle Circle) area() float64 {
	return 3.14 * circle.dir * circle.dir
}

//为了提高效率，通常我们方法和结构体的指针类型绑定
func (c2 *Circle) area2() float64 {

	fmt.Printf("c2 是  *Circle 指向的地址=%p\n", c2)

	c2.dir = 5
	return 3.14 * c2.dir * c2.dir
}

func main()  {
	var c Circle
	c.dir = 3
	area := c.area()
	fmt.Printf("%.2f\n",area)

	//编译器底层做了优化  (&c).area2() 等价 c.area()
	//因为编译器会自动的给加上 &c
	area2 := c.area2()
	fmt.Printf("%.2f\n",area2)
	fmt.Println("c.radius = ", c.dir) //10

}
