package main

import (
	"fmt"
)

//写入50个数，然后再读出来
func writeData(intChan chan int)  {
	for i:=0;i<50 ;i++  {
		intChan <- i
		fmt.Println("writeData ", i)
	}
	close(intChan)
}

func readData(exitChan chan bool,intChan chan int)  {
	//for val := range intChan {
	//	fmt.Println(val)
	//}

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}

	exitChan<-true
	close(exitChan)
}

func main()  {
	intChan := make(chan int,10)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	go readData(exitChan,intChan)

	for {
		_,ok := <-exitChan

		if !ok {
			break
		}
	}
}
