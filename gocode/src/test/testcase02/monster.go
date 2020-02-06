package testcase02

import (
"encoding/json"
"fmt"
"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func (m *Monster) Store()  bool{
	data,err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}
	filePath := "C:/code/go/gocode/src/test/testcase02/test.txt"
	err = ioutil.WriteFile(filePath,data,0666)
	if err != nil {
		fmt.Println("write file err =", err)
		return false
	}
	return true
}

//给Monster绑定方法ReStore, 可以将一个序列化的Monster,从文件中读取，
//并反序列化为Monster对象,检查反序列化，名字正确
func (this *Monster) ReStore() bool {

	//1. 先从文件中，读取序列化的字符串
	filePath := "C:/code/go/gocode/src/test/testcase02/test.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err =", err)
		return false
	}

	//2.使用读取到data []byte ,对反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("UnMarshal err =", err)
		return false
	}
	return true
}