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

	_,err = conn.Do("HMSet","User","name","qiping","age",30)
	if err != nil {
		fmt.Println("set err=",err)
		return
	}

	r,err := redis.Strings(conn.Do("HMGet","User","name","age"))
	if err != nil {
		fmt.Println("set err=",err)
		return
	}

	fmt.Printf("%T",r)

	for i,v := range r {
		//fmt.Printf("i = %v,v = %v \n",i,v)
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}