package main

import "fmt"

type Person10 struct {
	Name string
}

//函数
//对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
func test01(p10 Person10)  {
	fmt.Println(p10.Name)
}
func test02(p10 *Person10)  {
	fmt.Println(p10.Name)
}

//对于方法（如struct的方法），
//接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以
func (p10 Person10) test3()  {
	fmt.Println(p10.Name)
}

func (p10 *Person10) test4()  {
	fmt.Println(p10.Name)
}


func main()  {
	p := Person10{"tom"}
	test01(p)
	test02(&p)

	p.test3()
	fmt.Println("main() p.name=", p.Name) // tom

	(&p).test3() // 从形式上是传入地址，但是本质仍然是值拷贝

	fmt.Println("main() p.name=", p.Name) // tom

	(&p).test4()
	fmt.Println("main() p.name=", p.Name) // mary
	p.test4() // 等价 (&p).test04 , 从形式上是传入值类型，但是本质仍然是地址拷贝
}
