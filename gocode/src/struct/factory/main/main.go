package main

import (
	"fmt"
	"struct/factory/model"
)

func main()  {
	stu := model.GetNewStudent("qiping",29)

	fmt.Println(stu)
	fmt.Println(stu.GetNewAge())
}

