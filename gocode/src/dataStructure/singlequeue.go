package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理队列
type Queue struct {
	MaxSize int
	front   int    //指向队列 首部
	rear    int    //指向队列 尾部
	array   [5]int // 数组=>模拟队列
}

//添加数据到队列
func (queue *Queue) addQueue(val int) (err error) {
	if queue.rear == queue.MaxSize-1 {
		return errors.New("queue full")
	}
	queue.rear++ //指针后移
	queue.array[queue.rear] = val

	return
}

//从队列中弹出数据
func (queue *Queue) popQueue() (val int, err error) {
	if queue.front == queue.rear {
		return -1, errors.New("queue empty")
	}
	queue.front++
	val = queue.array[queue.front]

	return val, err
}

//显示队列, 找到队首，然后到遍历到队尾
//
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是:")
	//this.front 不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

func main() {
	queue := &Queue{
		MaxSize: 5,
		front:   -1,
		rear:    -1,
		array:   [5]int{},
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.addQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.popQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
