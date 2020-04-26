package main

import "fmt"

//定义一个HeroNode
type HeroNode2 struct {
	Num int
	Name string
	NickName string
	PreHero *HeroNode2
	NextHero *HeroNode2
}

//给链表末尾插入一个结点
func InsertLinkNode(head *HeroNode2,newNode *HeroNode2)  {
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
func InsertLinkNode2(head *HeroNode2,newNode *HeroNode2)  {
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

		if temp.NextHero != nil {
			temp.NextHero.PreHero = newNode
		}
		newNode.NextHero = temp.NextHero
		newNode.PreHero = temp
		temp.NextHero = newNode
	}
}

//删除节点
func deleteLinkNode(head *HeroNode2,num int)  {
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
		//temp.NextHero.NextHero.PreHero = temp
		temp.NextHero = temp.NextHero.NextHero

		if temp.NextHero != nil {
			temp.NextHero.PreHero = temp
		}

	}


}

//显示链表的所有结点信息
func listLinkNode(head *HeroNode2){
	if head.NextHero == nil {
		fmt.Println("链表空空如也")
		return
	}

	temp := head

	for{
		if temp.NextHero == nil {
			break
		}
		fmt.Printf("[%d %s %s] ==>",temp.NextHero.Num,temp.NextHero.Name,temp.NextHero.NickName)
		temp = temp.NextHero
	}
}

//显示链表的所有结点信息
func listLinkNode2(head *HeroNode2){
	if head.NextHero == nil {
		fmt.Println("链表空空如也")
		return
	}

	temp := head

	//2. 让temp定位到双向链表的最后结点
	for {
		if temp.NextHero == nil {
			break
		}
		temp = temp.NextHero
	}

	for{

		fmt.Printf("[%d %s %s] ==>",temp.Num,temp.Name,temp.NickName)
		temp = temp.PreHero

		if temp.PreHero == nil {
			break
		}
	}
}

func main()  {
	//创建一个头节点
	head := &HeroNode2{}

	heroNode1 := &HeroNode2{
		Num:      1,
		Name:     "松江",
		NickName: "及时雨",
		NextHero: nil,
	}
	heroNode2 := &HeroNode2{
		Num:      2,
		Name:     "林冲",
		NickName: "豹子头",
		NextHero: nil,
	}

	heroNode3 := &HeroNode2{
		Num:      6,
		Name:     "孙二娘",
		NickName: "毒蜘蛛",
		NextHero: nil,
	}
	heroNode4 := &HeroNode2{
		Num:      4,
		Name:     "鲁智深",
		NickName: "花和尚",
		NextHero: nil,
	}

	InsertLinkNode2(head,heroNode1)
	InsertLinkNode2(head,heroNode2)
	InsertLinkNode2(head,heroNode3)
	InsertLinkNode2(head,heroNode4)

	listLinkNode(head)
	fmt.Println()
	listLinkNode2(head)

	deleteLinkNode(head,2)
	fmt.Println()
	listLinkNode2(head)
}
