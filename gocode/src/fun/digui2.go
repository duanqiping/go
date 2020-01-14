package main

import "fmt"

//题3：猴子吃桃子问题有一堆桃子，猴子第一天吃了其中的一半，
//并再多吃了一个！以后每天猴子都吃其中的一半，然后再多吃一个。
//当到第十天时，想再吃时（还没吃），发现只有1个桃子了。问题：最初共多少个桃子？

// 10  	1
//9  	(1+1)*2 = 4  第9天有几个桃子  =  (第10天桃子数量 + 1) * 2
//8       (4+1)*2 = 10  peach(n) = (peach(n+1) + 1) * 2

func peach(n int) int {
	if n > 10 || n < 1{
		fmt.Printf("你的输入有误\n")
		return 0
	}
	if n == 10{
		return 1
	}else{
		return (peach(n+1)+1)*2
	}
}

func main()  {
	n := 1
	fmt.Printf("第%v 还有 %v 个桃子\n",n,peach(n))
}