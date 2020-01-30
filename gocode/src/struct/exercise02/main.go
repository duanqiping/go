package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//1.声明Hero结构体
type Hero struct {
	Name string
	Age int
}

//2.声明一个Hero结构体切片类型
type HeroSlice []Hero


func (h HeroSlice) Len() int {
	return len(h)
}

//Less方法就是决定你使用什么标准进行排序
//1. 按Hero的年龄从小到大排序!
func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}
func (h HeroSlice) Swap(i, j int) {
	//交换
	//temp := h[i]
	//h[i] = h[j]
	//h[j] = temp
	//下面的一句话等价于三句话
	h[i],h[j] = h[j],h[i]
}

func main()  {
	//先定义一个数组/切片
	var sliceSort = []int{1,2,0,-19,100}

	//Ints函数将a排序为递增顺序。
	sort.Ints(sliceSort)

	fmt.Println(sliceSort)

	//请大家对结构体切片进行排序
	//1. 冒泡排序...
	//2. 也可以使用系统提供的方法

	//测试看看我们是否可以对结构体切片进行排序
	var heros HeroSlice
	for i := 0 ; i < 10; i++ {
		hero := Hero{
			Name:fmt.Sprintf("英雄 | %d",rand.Intn(100)),
			Age:rand.Intn(100),
		}
		//将 hero append到 heroes切片
		heros = append(heros,hero)
	}
	//看看排序前的顺序
	for _ , v := range heros {
		fmt.Println(v)
	}

	//调用sort.Sort
	sort.Sort(heros)
	fmt.Println("-----------排序后------------")
	//看看排序后的顺序
	for _ , v := range heros {
		fmt.Println(v)
	}

	i := 10
	j := 20
	i, j = j, i
	fmt.Println("i=", i, "j=", j) // i=20 j = 10
}
