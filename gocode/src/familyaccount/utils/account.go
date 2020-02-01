package utils

import "fmt"

type account struct {
	name string
	pwd string
}

func NewAccount() *account {
	return &account{}
}

//验证账号
func (a account) Verify() {
	for {
		fmt.Printf("请输入账号：")
		fmt.Scanln(&a.name)
		fmt.Printf("请输入密码：")
		fmt.Scanln(&a.pwd)
		if a.name == "qiping" && a.pwd == "111222"{
			break
		}else{
			fmt.Printf("请输入账号或密码有误，请重新输入\n")
		}
	}
}


//这种写法 不需要强转
type intger = int
func test()  {
	var a intger
	a = 1
	
	var b int
	
	b = a
	
	fmt.Println(b)
	
	
	
}
