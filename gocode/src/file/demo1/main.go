package main

import (
	"fmt"
	"os"
)

//打开文件
//概念说明: file 的叫法
//1. file 叫 file对象
//2. file 叫 file指针
//3. file 叫 file 文件句柄

func main()  {

	filePath := "C:/code/go/test.txt"
	file,err := os.Open(filePath)

	if err != nil {
		fmt.Println("open file err=", err)
	}

	fmt.Printf("%T",file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}


}
