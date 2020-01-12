package main

//包的引入默认 从src定位的
import (
	"fmt"
	u "utls"
)

func main()  {
	u.ReturnHelloWorld()
	fmt.Println("num =",u.Num1)
}
