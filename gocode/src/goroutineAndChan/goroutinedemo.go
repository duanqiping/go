package main

import (
	"fmt"
	"time"
)

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔1秒输出 "hello,world"
// 在主线程中也每隔一秒输出"hello,golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行

//编写一个函数，每隔1秒输出 "hello,world"
func test()  {
	for i := 0;i<10 ;i++  {
		fmt.Println("hello,golang")
		time.Sleep(time.Second)
	}
}

func main()  {
	go test()

	for i := 0;i<10 ;i++  {
		fmt.Println("main() hello,golang")
		time.Sleep(time.Second)
	}
}
