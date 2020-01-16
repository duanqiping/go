package main

import (
	"fmt"
	"time"
)

func main()  {
	//看看日期和时间相关函数和方法使用
	//1. 获取当前时间

	now := time.Now()

	fmt.Println("当前时间为：",now)
	fmt.Println("当前时间 年为：",now.Year())//年
	fmt.Println("当前时间 月为：",now.Month())//月
	fmt.Println("当前时间 月为：",int(now.Month()))//月
	fmt.Println("当前时间 日为：",now.Day())//日

	fmt.Println("当前时间 为时：",now.Hour())//时
	fmt.Println("当前时间 分为：",now.Minute())//分
	fmt.Println("当前时间 秒为：",now.Second())//秒

	//格式化日期时间

	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dateStr=%v\n", dateStr)

	//格式化日期时间的第二种方式
	fmt.Println("时间",now.Format("2006-01-02 15:04:05"))
	fmt.Println("时间",now.Format("2006-01-02"))
	fmt.Println("时间",now.Format("15:04:05"))

}
