package main

import (
	"fmt"       // 导入 fmt 包，打印字符串是需要用到
	"strconv"
)

const (
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

// 将NewInt定义为int类型
type NewInt int
// 将int取一个别名叫IntAlias
type IntAlias = int

func main()  {

	i, _ := strconv.Atoi("3")
	fmt.Println(3 + i) // 6
	// Atoi()转换失败
	i, err := strconv.Atoi("a")
	fmt.Println(err)
	if err != nil {
		fmt.Println("converted failed")
	}

	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)

	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)//十进制格式输出
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)//二进制格式输出

	// 输出所有枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)

	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)
}

type Weapon int

const (
	Arrow Weapon = iota    // 开始生成枚举值, 默认为0
	Shuriken
	SniperRifle
	Rifle
	Blower
)

