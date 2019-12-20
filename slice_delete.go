package main

import "fmt"

func main() {

	seq := []string{"a", "b", "c", "d", "e"}
	// 指定删除位置
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)


	var a = []int{1, 2, 3}

	//a = a[1:] // 删除开头1个元素
	//a = a[1:] // 删除开头1个元素

	//a = append(a[:0], a[1:]...) // 删除开头1个元素
	//a = append(a[:0], a[2:]...) // 删除开头N个元素

	//var c = copy(a, a[1:])

	//a = a[:copy(a, a[1:])] // 删除开头1个元素
	//a = a[:copy(a, a[N:])] // 删除开头N个元素

	//a = a[:2] // 删除开头1个元素

	var b = []int{1, 2, 3}
	b = append(b[:0], b[2:]...) // 删除开头1个元素
	//b = append(b[:0], b[N:]...) // 删除开头N个元素


	//var a = []int{1, 2, 3,4,5,6}
	//a = append(a[:3], a[3+1:]...) // 删除中间1个元素
	//a = append(a[:3], a[3+3:]...) // 删除中间N个元素
	//a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	//a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

	//a = a[:len(a)-1] // 删除尾部1个元素
	a = a[:len(a)-2] // 删除尾部N个元素

	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(c)

}