package main

import "fmt"

type Cat struct {
	Num int
	Name string
	NextCat *Cat
}

func insertCat(head *Cat,NewCat *Cat)  {
	if head.NextCat == nil {

		head.Num = NewCat.Num
		head.Name = NewCat.Name
		head.NextCat = head

		fmt.Println(NewCat,"加入到环形链表")
		return
	}
	//定义一个临时变量，帮忙,找到环形的最后结点
	temp := head
	for {
		if temp.NextCat == head {
			break
		}
		temp = temp.NextCat
	}
	temp.NextCat = NewCat
	NewCat.NextCat = head
}

func ListLink(head *Cat)  {
	fmt.Println("环形链表的情况如下：")
	if head.NextCat == nil {
		fmt.Println("环形链表为空")
		return
	}
	temp := head
	for {

		fmt.Printf("cat 为 %v \n",temp)
		if temp.NextCat == head {
			break
		}

		temp = temp.NextCat
	}
}

//删除一个环形单向链表的思路如下：
//1. 先让temp 指向 head
//2. 让 helper 指向环形链表最后.
//3. 让temp  和 要删除的id 进行比较，如果相同，则同 helepr完成删除[这里必须考虑如果删除就是头结点]
func DeleteCat(head *Cat,num int)  *Cat {
	temp := head
	helper := head

	if temp.NextCat == nil {
		fmt.Println("链表空空如也")
		return head
	}

	if temp.NextCat == head {
		if temp.Num == num {
			head.NextCat = nil
			fmt.Println("只有一个节点，删除成功")
			return head
		}
	}

	//将helper 定位到链表最后
	for {
		if helper.NextCat == head {
			break
		}
		helper = helper.NextCat
	}

	//如果有两个包含两个以上结点
	flag := true
	for {
		if temp.NextCat == head { //如果到这来，说明我比较到最后一个【最后一个还没比较】
			break
		}
		if temp.Num == num {
			if temp == head { //说明删除的是头结点
				head = head.NextCat
			}
			//恭喜找到., 我们也可以在直接删除
			helper.NextCat = temp.NextCat
			fmt.Printf("猫猫=%d\n", num)
			flag = false
			break
		}
		temp = temp.NextCat //移动 【比较】
		helper = helper.NextCat //移动 【一旦找到要删除的结点 helper】
	}

	//这里还有比较一次
	if flag { //如果flag 为真，则我们上面没有删除
		if temp.Num == num {
			helper.NextCat = temp.NextCat
			fmt.Printf("猫猫=%d\n", num)
		}else {
			fmt.Printf("对不起，没有no=%d\n", num)
		}
	}
	return head

}



func main()  {
	head := &Cat{}
	cat1 := &Cat{
		Num:     1,
		Name:    "tom",
		NextCat: nil,
	}
	cat2 := &Cat{
		Num:     2,
		Name:    "jack",
		NextCat: nil,
	}

	cat3 := &Cat{
		Num:     3,
		Name:    "jack2",
		NextCat: nil,
	}
	insertCat(head,cat1)
	insertCat(head,cat2)
	insertCat(head,cat3)
	ListLink(head)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	head = DeleteCat(head,2)
	head = DeleteCat(head,10)
	head = DeleteCat(head,1)
	ListLink(head)

}
