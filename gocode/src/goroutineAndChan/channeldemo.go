package main

import "fmt"

func main()  {
	//演示一下管道的使用
	//1. 创建一个可以存放3个int类型的管道

	var myChan chan int
	myChan = make(chan int,3)

	//2. 看看intChan是什么
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", myChan, &myChan)

	myChan <- 1
	myChan <- 2
	myChan <- 3
	<-myChan
	myChan <- 4

	//4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(myChan), cap(myChan)) // 3, 3

	var num2 int
	num2 = <-myChan
	fmt.Println("num2=", num2)

	num3 := <-myChan
	num4 := <-myChan

	//num5 := <-myChan

	fmt.Println("num3=", num3, "num4=", num4/*, "num5=", num5*/)
	//fmt.Println("num3=", num3, "num4=", num4, "num5=", num5)
}
