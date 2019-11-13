package main

import "container/list"
import "fmt"

func main() {
	l := list.New()

	// 尾部添加
	l.PushBack("canon")

	// 头部添加
	l.PushFront(67)

	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")

	// 在fist之后添加high
	l.InsertAfter("high", element)

	// 在fist之前添加noon
	l.InsertBefore("noon", element)

	fmt.Println(element)

	// 使用
	l.Remove(element)

	l2 := list.New()

	// 尾部添加
	l2.PushBack("canon")

	// 头部添加
	l2.PushFront(67)

	for i := l2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}