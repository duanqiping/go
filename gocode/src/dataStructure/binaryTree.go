package main

import "fmt"

type  Hero struct {
	num int
	name string
	leftNode *Hero
	rightNode *Hero
}
//二叉树前序遍历： 访问根节点->前序遍历左子树->前序遍历右子树
func PreOrder(Node *Hero)  {
	if Node != nil {
		fmt.Printf("序号： %d, 名称为：%s \n", Node.num,Node.name)
		PreOrder(Node.leftNode)
		PreOrder(Node.rightNode)
	}
}

//二叉树中序遍历： 中序遍历左子树->访问根节点->中序遍历右子树
func InfixOrder(Node *Hero)  {
	if Node != nil {
		PreOrder(Node.leftNode)
		fmt.Printf("序号： %d, 名称为：%s \n", Node.num,Node.name)
		PreOrder(Node.rightNode)
	}
}

//二叉树后序遍历： 后序遍历左子树->后序遍历右子树->访问根节点
func PostOrder(Node *Hero)  {
	if Node != nil {
		PreOrder(Node.leftNode)
		PreOrder(Node.rightNode)
		fmt.Printf("序号： %d, 名称为：%s \n", Node.num,Node.name)
	}
}

func main()  {

	rootHero := Hero{
		num:      1,
		name:     "宋江",
		leftNode: nil,
		rightNode:  nil,
	}

	left1 := Hero{
		num:      2,
		name:     "吴用",
		leftNode: nil,
		rightNode:    nil,
	}
	right1 := Hero{
		num:      3,
		name:     "卢俊义",
		leftNode: nil,
		rightNode:    nil,
	}

	left10 := Hero{
		num:      10,
		name:     "jack10",
		leftNode: nil,
		rightNode:    nil,
	}

	right12 := Hero{
		num:      12,
		name:     "tom12",
		leftNode: nil,
		rightNode:    nil,
	}

	right_r_2 := Hero{
		num:       4,
		name:      "林冲",
		leftNode:  nil,
		rightNode: nil,
	}

	rootHero.leftNode = &left1
	rootHero.rightNode = &right1
	left1.leftNode = &left10
	left1.rightNode = &right12
	right1.rightNode = &right_r_2

	//PreOrder(&rootHero)
	//InfixOrder(&rootHero)
	PostOrder(&rootHero)





}
