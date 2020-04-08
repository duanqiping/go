package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int // 4
	array [5]int // 数组
	head  int //指向队列队首 0
	tail int  //指向队尾 0
}

func (queue *CircleQueue) IsFull()  bool{
	return (queue.tail + 1)%queue.maxSize == queue.head
}

func (queue *CircleQueue) Push(val int)  (err error){
	if queue.IsFull() {
		return errors.New("queue full")
	}
	//分析出this.tail 在队列尾部，但是包含最后的元素
	queue.array[queue.tail] = val
	queue.tail = (queue.tail + 1)%queue.maxSize

	return
}

func (queue *CircleQueue) IsEmpty()  bool{
	return (queue.head + 1)%queue.maxSize == queue.tail
}
func (queue *CircleQueue) Pop()  (val int,err error)  {
	if queue.IsEmpty() {
		return -1,errors.New("is empty")
	}
	val = queue.array[queue.head]
	queue.head = (queue.head + 1)%queue.maxSize
	return val,err
}

//取出环形队列有多少个元素
func (queue *CircleQueue) Size() int {
	//这是一个关键算法
	return (queue.tail - queue.head + queue.maxSize) % queue.maxSize

	//return len(queue.array) //数组初始化长度就为4
}

//显示队列
func (queue *CircleQueue) ListQueue()  {
	fmt.Println("环形队列情况如下：")

	size := queue.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	//设计一个辅助的变量，指向head
	tempHead := queue.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, queue.array[tempHead])
		tempHead = (tempHead + 1) % queue.maxSize
	}
	fmt.Println()
}


func main()  {
	//初始化一个环形队列
	queue := &CircleQueue{
		maxSize : 5,
		head : 0,
		tail : 0,
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
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}