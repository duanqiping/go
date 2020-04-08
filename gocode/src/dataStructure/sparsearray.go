package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//1.	应用实例：保存二维数组（类似棋盘、游戏地图等）
//2.	把稀疏数组存盘，并且可以恢复原来的二维数组

type ValNode struct {
	row int
	col int
	val int
}

func writeFile(str string)  {
	filePath := "C:/code/go/sparsearray.txt"

	file,err := os.OpenFile(filePath,os.O_APPEND | os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("文件创建失败 %v",err)
		return
	}
	//及时关闭file句柄
	defer file.Close()

	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	writer.WriteString(str)

	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()
}

func main()  {
	var chessMap [11][11] int
	chessMap [1][2] = 1 //黑子
	chessMap [5][6] = 2 //白子

	fmt.Println("原始数组")
	for i := 0;i<11 ; i++ {
		for j := 0; j < 11;j++  {
			fmt.Printf("%d\t",chessMap[i][j])
		}
		fmt.Println()
	}

	//转成稀疏数组 思路
	//遍历数组 如果值不为0，则创建一个varNode结构体
	//然后将结构体 赋值给切片
	var sparseArr []ValNode
	varNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr,varNode)

	for i := 0;i<11 ; i++ {
		for j := 0; j < 11;j++  {
			if chessMap[i][j] != 0 {
				varNode := ValNode{
					row: i,
					col: j,
					val: chessMap[i][j],
				}
				sparseArr = append(sparseArr,varNode)
			}
		}
	}

	fmt.Println("遍历稀疏数组")
	for i := 0;i< len(sparseArr); i++ {
		fmt.Printf("%d\t%d\t%d\t",sparseArr[i].row,sparseArr[i].col,sparseArr[i].val)
		fmt.Println()
	}

	//将稀疏数组存盘
	for i := 0;i< len(sparseArr); i++ {

		str := strconv.Itoa(sparseArr[i].row) + " " + strconv.Itoa(sparseArr[i].col) + " " + strconv.Itoa(sparseArr[i].val)
		str += "\r\n"

		fmt.Println(str)

		writeFile(str)
	}


	//将稀疏数组复原
	var chessMap2 [11][11]int
	for i := 0;i< len(sparseArr); i++ {
		if i != 0 {
			chessMap2[sparseArr[i].row][sparseArr[i].col] = sparseArr[i].val
		}
	}

	fmt.Println("遍历恢复稀疏数组后的数组")
	for i := 0;i<11 ; i++ {
		for j := 0; j < 11;j++  {
			fmt.Printf("%d\t",chessMap2[i][j])
		}
		fmt.Println()
	}

}
