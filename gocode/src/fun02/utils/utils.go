package utils

import "fmt"

var Age = 10
var Name = "tom"

func init()  {
	fmt.Println("引入包的init 函数")
	Age = 100
	Name = "qiping"
}

func Test01()  {
	fmt.Println("Hello world")
}
