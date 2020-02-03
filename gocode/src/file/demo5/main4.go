package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//自己编写一个函数，接收两个文件路径 srcFileName(原文件路径) dstFileName(要copy得到的文件路径)
func copyFile(dstFileName string,srcFileName string) (written int64, err error)  {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()

	//通过srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)

}

func main()  {
	//将d:/flower.jpg 文件拷贝到 e:/abc.jpg
	//调用CopyFile 完成文件拷贝

	//txt文件 复制后的文件 是空的
	//filePath1 := "C:/code/go/abc.txt"
	//filePath2 := "C:/code/go/nnn.txt"

	filePath1 := "C:/code/go/aa.jpg"
	filePath2 := "C:/code/go/bb.jpg"

	_,err := copyFile(filePath2,filePath1)

	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}



}
