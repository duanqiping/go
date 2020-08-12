package main

import "fmt"

type child struct {
	num int
	nextChild *child
}

func addChild(count int) *child {
	first := &child{}
	current := &child{}

	if count < 1 {
		fmt.Println("玩游戏的小孩数里不能少于1")
		return first
	}

	for i:=1;i<=count;i++{
		child := &child{
			num:       i,
			nextChild: nil,
		}
		if i == 1 {
			first = child
			current = first
			current.nextChild = first
		}else{
			current.nextChild = child
			current = child
			current.nextChild = first
		}
	}

	return first
}

func showChild(first *child)  {
	if first.nextChild == nil {
		fmt.Println("玩游戏的小孩不能为空")
		return
	}

	temp := first
	for {
		fmt.Printf("小孩编号：%d\n",temp.num)
		if temp.nextChild == first {
			break
		}
		temp = temp.nextChild
	}
}

//约瑟夫问题练习
/*
设编号为1，2，… n的n个人围坐一圈，约定编号为k（1<=k<=n）
的人从1开始报数，数到m 的那个人出列，它的下一位又从1开始报数，
数到m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
*/
func childPlayGame(first *child,startN int,offset int)  {
	if first.nextChild == nil {
		fmt.Println("玩游戏的小孩不能为空")
		return
	}

	tail := first
	for {
		if tail.nextChild == first {
			break
		}
		tail = tail.nextChild
	}

	for i := 1; i<=startN-1; i++ {
		tail = tail.nextChild
		first = first.nextChild
	}

	fmt.Println()
	for {

		if first.nextChild == tail.nextChild {
			break
		}

		for i:=1;i<=offset-1;i++{
			tail = tail.nextChild
			first = first.nextChild
		}

		fmt.Printf("编号为%d出列\n",first.num)
		tail.nextChild = first.nextChild
		first = first.nextChild
	}
	fmt.Printf("编号为%d出列\n",first.num)

}

func main()  {
	first := addChild(6)
	showChild(first)
	childPlayGame(first,2,3)

}
