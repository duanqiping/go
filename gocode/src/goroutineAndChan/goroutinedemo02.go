package main

import (
	"fmt"
	"sync"
)

// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
// 最后显示出来。要求使用goroutine完成

// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
// 2. 我们启动的协程多个，统计的将结果放入到 map中
// 3. map 应该做出一个全局的.

var (
	intMap = make(map[int]int)

	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁，
	//sync 是包: synchornized 同步
	//Mutex : 是互斥
	lock sync.Mutex

)

func a(n int)  {
	sum := 1
	for i := 1;i<=n ;i++  {
		sum *= i
	}
	lock.Lock()
	intMap[n] = sum
	lock.Unlock()
}

func main()  {

	for i := 1;i<=20;i++ {
		go a(i)
	}

	for index,v := range intMap{
		fmt.Printf("%v = %v \n",index,v)
	}
	//经测试 可以输出map  和教程 的结果 有出入


	//休眠10秒钟【第二个问题 】
	//time.Sleep(time.Second * 5)

	//这里我们输出结果,变量这个结果
	lock.Lock()
	for i, v := range intMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()

	
}
