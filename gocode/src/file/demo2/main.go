package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	//打开文件
	//概念说明: file 的叫法
	//1. file 叫 file对象
	//2. file 叫 file指针
	//3. file 叫 file 文件句柄
	filePath := "C:/code/go/test.txt"
	file , err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err=", err)
	}

	defer file.Close()

	// 创建一个 *Reader  ，是带缓冲的
	/*
		const (
		defaultBufSize = 4096 //默认的缓冲区为4096
		)
	*/
	reader := bufio.NewReader(file)
	fmt.Printf("%T\n",reader)

	//for {
	//  //这个方法有个缺陷，如果最后一行没有换号，有可能会把最后一行字符串漏掉
	//	fileString,err := reader.ReadString('\n')
	//	if err == io.EOF { // io.EOF表示文件的末尾
	//		break
	//	}
	//	//输出内容
	//	fmt.Printf(fileString)
	//}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}


	fmt.Println("文件读取结束...")

}
