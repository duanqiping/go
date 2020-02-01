package main

import "fmt"

//声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

//让Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

func (p Phone) Call() {
	fmt.Println("手机 在打电话..")
}

type Camera struct {
	name string
}
//让Camera 实现   Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

type Computer struct {

}

func (c Computer) working(usb Usb)  {
	usb.Start()
	//如果usb是指向Phone结构体变量，则还需要调用Call方法
	//类型断言..[注意体会!!!]
	if Phone,ok := usb.(Phone); ok {
		Phone.Call()
	}
	usb.Stop()
}

func main()  {
	phone := Phone{}
	camera := Camera{}

	c := Computer{}
	c.working(phone)
	c.working(camera)
}
