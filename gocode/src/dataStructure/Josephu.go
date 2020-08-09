package main

import "fmt"

//约瑟夫丢手帕问题

type Boy struct {
	num int
	nextBoy *Boy
}

//分析思路
//1. 编写一个函数，PlayGame(first *Boy, startNo int, countNum int)
//2. 最后我们使用一个算法，按照要求，在环形链表中留下最后一个人
func PlayGame(first *Boy, startNo int, countNum int) {

	//1. 空的链表我们单独的处理
	if first.nextBoy == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	//留一个，判断 startNO <= 小孩的总数
	//2. 需要定义辅助指针，帮助我们删除小孩
	tail := first
	//3. 让tail执行环形链表的最后一个小孩,这个非常的重要
	//因为tail 在删除小孩时需要使用到.
	for {
		if tail.nextBoy == first { //说明tail到了最后的小孩
			break
		}
		tail = tail.nextBoy
	}
	//4. 让first 移动到 startNo [后面我们删除小孩，就以first为准]
	for i := 1; i <= startNo - 1; i++ {
		first = first.nextBoy
		tail = tail.nextBoy
	}
	fmt.Println()
	//5. 开始数 countNum, 然后就删除first 指向的小孩
	for {
		//开始数countNum-1次
		for i := 1; i <= countNum -1; i++ {
			first = first.nextBoy
			tail = tail.nextBoy
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.num)
		//删除first执行的小孩
		first = first.nextBoy
		tail.nextBoy = first
		//判断如果 tail == first, 圈子中只有一个小孩.
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩小孩编号为%d 出圈 \n", first.nextBoy)

}
func AddBoy(num int) *Boy {

	first := &Boy{} //空结点
	curBoy := &Boy{} //空结点

	//判断
	if num < 1 	{
		fmt.Println("num的值不对")
		return first
	}
	//循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			num : i,
		}
		//分析构成循环链表，需要一个辅助指针[帮忙的]
		//1. 因为第一个小孩比较特殊
		if i == 1 { //第一个小孩
			first = boy //不要动
			curBoy = boy
			curBoy.nextBoy = first //
		} else {
			curBoy.nextBoy = boy
			curBoy = boy
			curBoy.nextBoy = first //构造环形链表
		}
	}
	return first

}

func ShowBoy(head *Boy)  {
	temp := head
	if temp.nextBoy == nil {
		fmt.Println("老兄,一个人都还没有呀")
		return
	}

	for {
		if temp.nextBoy == head {
			break
		}

		fmt.Printf("编号 = %d\n",temp.num)
		temp = temp.nextBoy
	}
}


/*
设编号为1，2，… n的n个人围坐一圈，约定编号为k（1<=k<=n）
的人从1开始报数，数到m 的那个人出列，它的下一位又从1开始报数，
数到m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
*/

//分析思路
//1. 编写一个函数，PlayGame(first *Boy, startNo int, countNum int)
//2. 最后我们使用一个算法，按照要求，在环形链表中留下最后一个人
func main()  {

	head := AddBoy(5)

	ShowBoy(head)

	PlayGame(head, 1, 3)
}
