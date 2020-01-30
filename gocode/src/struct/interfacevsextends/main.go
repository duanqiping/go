package main

import "fmt"

type Monkey struct {
	Name string
}

//声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (monkey Monkey) Climbing()  {
	fmt.Println(monkey.Name,"生来会爬树....")
}

//LittleMonkey结构体
type LittleMonkey struct {
	Monkey
}

func (mon LittleMonkey) Flying()  {
	fmt.Println(mon.Name,"通过学习可以飞...")
}

func (mon LittleMonkey) Swimming()  {
	fmt.Println(mon.Name,"通过学习可以游泳...")
}


func main()  {
	monkey := LittleMonkey{Monkey{"孙悟空"}}
	monkey.Climbing()

	monkey.Flying()
	monkey.Swimming()


}
