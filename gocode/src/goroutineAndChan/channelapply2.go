package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int,Num int)  {
	for i := 1; i <= Num; i++ {
		intChan<- i
	}
	close(intChan)
}

// 从 intChan取出数据，并判断是否为素数,如果是，就
// 	//放入到primeChan
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool)  {
	for{
		num,ok := <-intChan
		if !ok {
			break
		}

		//判断是否为素数
		flag := true
		for i := 2;i<num ;i++  {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan<-num
		}


	}
	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	//这里我们还不能关闭 primeChan
	//向 exitChan 写入true
	exitChan<- true
}

func main()  {
	//目标 统计1到8000里面有多少个素数 putNumChan
	//创建 8 个协程 获取素数 存入管道 storeNumChan
	//创建一个退出管道 保存那8个协程 的退出状态 exitChan

	pNum := 2 //要创建的协程个数
	Num := 100000 //从多大的数里面获取素数

	intChan := make(chan int,1000)
	primeChan := make(chan int,10000)
	exitChan := make(chan bool,4)

	start := time.Now().Unix()

	go putNum(intChan,Num)

	for i:=0; i<pNum;i++  {
		go primeNum(intChan,primeChan,exitChan)
	}

	go func() {
		for i:=0;i<pNum ;i++  {
			<- exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end - start)

		close(primeChan)
	}()

	for {
		_,ok := <- primeChan
		if !ok {
			break
		}
		//将结果输出
		//fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main线程退出")

}
