package main

import "fmt"

type girl struct {
	num int
	nextGirl *girl
}

//添加玩游戏的女孩个数，并返回第一个女孩
func addGirl(number int)  *girl{

	first := &girl{}
	cur := &girl{}

	if number < 1 {
		fmt.Println("玩游戏的女孩个数不能小于1")
		return first
	}

	for i:=1;i<=number;i++ {

		girl := &girl{
			num:      i,
			nextGirl: nil,
		}

		if i == 1 {
			first = girl
			cur = first
			cur.nextGirl = first
		}else{
			cur.nextGirl = girl
			cur = girl
			cur.nextGirl = first
		}
	}
	return first
}

func showGirl(first *girl)  {
	if first.nextGirl == nil {
		fmt.Println("玩游戏的女孩个数不能小于1")
		return
	}

	cur := first
	for {
		fmt.Printf("玩游戏的女孩编号为：%d\n",cur.num)

		if cur.nextGirl == first {
			return
		}
		cur = cur.nextGirl
	}
}

func girlPlayGame(first *girl,startN int,offset int)  {

	tail := first

	//startN 不能大于 总数
	for {
		if tail.nextGirl == first {
			break
		}
		tail = tail.nextGirl
	}

	for i := 1;i<=startN-1 ;i++  {
		first = first.nextGirl
		tail = tail.nextGirl
	}

	for{
		for i := 1;i<=offset-1;i++  {
			first = first.nextGirl
			tail = tail.nextGirl
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.num)
		//删除first执行的小孩
		first = first.nextGirl
		tail.nextGirl = first
		//判断如果 tail == first, 圈子中只有一个小孩.
		if tail == first {
			break
		}
	}
	fmt.Printf("最后一个小孩编号为%d 出圈 \n", first.num)


}

//约瑟夫问题练习
/*
设编号为1，2，… n的n个人围坐一圈，约定编号为k（1<=k<=n）
的人从1开始报数，数到m 的那个人出列，它的下一位又从1开始报数，
数到m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
*/
func main()  {
	first := addGirl(10)
	showGirl(first)

	girlPlayGame(first,2,3)

}