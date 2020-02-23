package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	//通过go 向redis 写入数据和读取数据
	//1. 链接到redis
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=",err)
		return
	}
	defer conn.Close()//关闭

	_,err = conn.Do("HSet","User","name","qiping")
	if err != nil {
		fmt.Println("set err=",err)
		return
	}

	_,err = conn.Do("HSet","User","age",29)
	if err != nil {
		fmt.Println("set err=",err)
		return
	}

	r,err := redis.String(conn.Do("HGet","User","name"))
	if err != nil {
		fmt.Println("set err=",err)
		return
	}

	age,err := redis.Int(conn.Do("HGet","User","age"))
	if err != nil {
		fmt.Println("set err=",err)
		return
	}

	fmt.Printf("操作ok 名字 %v,年龄 %v",r,age)
}
