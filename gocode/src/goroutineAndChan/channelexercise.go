package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

func main()  {
	allChan := make(chan interface{},10)

	allChan<-  10
	allChan<-  "tom jack"
	cat := Cat{"小花猫",4}
	allChan <- cat

	//我们希望获得到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan

	newCat := <-allChan //从管道中取出的Cat是什么?
	fmt.Printf("newCat = %v",newCat)

	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)
	//下面的写法是错误的!编译不通过
	//fmt.Printf("newCat.Name=%v", newCat.Name)
	//使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)
}
