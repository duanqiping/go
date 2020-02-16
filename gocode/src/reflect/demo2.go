package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var str = "tom"   //ok
	fs := reflect.ValueOf(&str) //ok fs -> string

	fmt.Printf("类型是：%T  值是：%v\n",fs,fs)

	fs.Elem().SetString("jack")
	//*fs = "jack";
	fmt.Println(str)
}
