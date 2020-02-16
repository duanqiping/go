package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn)  {

	//这里我们循环的接收客户端发送的数据
	defer conn.Close()

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1. 等待客户端通过conn发送信息
		//2. 如果客户端没有wrtie[发送]，那么协程就阻塞在这里
		//fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Printf("客户端退出 err=%v", err)
			return //!!!
		}
		//3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}


func main()  {
	fmt.Println("服务器开始监听....")

	//net.Listen("tcp", "0.0.0.0:8888")
	//1. tcp 表示使用网络协议是tcp
	//2. 0.0.0.0:8888 表示在本地监听 8888端口
	Listener, err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	fmt.Println("监听成功",Listener)

	defer Listener.Close()

	for {
		//等待客户端链接
		fmt.Println("等待客户端来链接....")
		conn,err := Listener.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		}else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}

		go process(conn)
	}
}