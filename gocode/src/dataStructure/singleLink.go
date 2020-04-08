package main

import "fmt"

//定义一个HeroNode
type HeroNode struct {
	Num int
	Name string
	NickName string
	NextHero *HeroNode
}

//给链表末尾插入一个结点
func InsertNode(head *HeroNode,newNode *HeroNode)  {
	//思路
	//1. 先找到该链表的最后这个结点
	//2. 创建一个辅助结点[跑龙套, 帮忙]
	temp := head
	for {
		if temp.NextHero == nil {
			//表示找到最后
			break
		}
		//让temp不断指向下一个节点
		temp = temp.NextHero
	}
	//将新节点挂到链表末节点
	temp.NextHero = newNode
}

//根据 num 大小插入一个新的节点 （实用）
func InsertNode2(head *HeroNode,newNode *HeroNode)  {
	temp := head
	flag := false
	for {
		if temp.NextHero == nil {
			//表示找到最后
			break
		}else if temp.NextHero.Num > newNode.Num {
			//说明newHeroNode 就应该插入到temp后面
			break
		}else if temp.NextHero.Num == newNode.Num {
			flag = true
			break
		}

		//让temp不断指向下一个节点
		temp = temp.NextHero
	}
	if flag {
		fmt.Println("节点已经存在：",newNode.Name)
		return
	}else {
		newNode.NextHero = temp.NextHero
		temp.NextHero = newNode
	}
}

//删除节点
func deleteNode(head *HeroNode,num int)  {
	temp := head
	flag := false
	for{
		if temp.NextHero == nil {
			break
		}else if temp.NextHero.Num == num {
			//找到要删除的节点
			flag = true
			break
		}
		temp = temp.NextHero
	}
	if flag == false {
		fmt.Println("要删除的节点不存在：",num)
	}else{
		temp.NextHero = temp.NextHero.NextHero
	}


}

//显示链表的所有结点信息
func listHero(head *HeroNode){
	if head.NextHero == nil {
		fmt.Println("链表空空如也")
		return
	}

	temp := head

	for{
		if temp.NextHero == nil {
			return
		}
		fmt.Printf("[%d %s %s] ==>",temp.NextHero.Num,temp.NextHero.Name,temp.NextHero.NickName)
		temp = temp.NextHero
	}
}

func main()  {
	//创建一个头节点
	head := &HeroNode{}

	heroNode1 := &HeroNode{
		Num:      1,
		Name:     "松江",
		NickName: "及时雨",
		NextHero: nil,
	}
	heroNode2 := &HeroNode{
		Num:      2,
		Name:     "林冲",
		NickName: "豹子头",
		NextHero: nil,
	}

	heroNode3 := &HeroNode{
		Num:      6,
		Name:     "孙二娘",
		NickName: "毒蜘蛛",
		NextHero: nil,
	}
	heroNode4 := &HeroNode{
		Num:      4,
		Name:     "鲁智深",
		NickName: "花和尚",
		NextHero: nil,
	}

	InsertNode2(head,heroNode1)
	InsertNode2(head,heroNode2)
	InsertNode2(head,heroNode3)
	InsertNode2(head,heroNode4)

	deleteNode(head,5)

	listHero(head)
}
