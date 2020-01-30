package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

//type Usb2 interface {
//	Start()
//	Stop()
//	test()
//}

type Phone struct {

}

func (p Phone) Start()  {
	fmt.Println("phone 开始工作...")
}
func (p Phone) Stop()  {
	fmt.Println("phone 停止工作...")
}

type Camera struct {

}

func (c Camera) Start()  {
	fmt.Println("相机 开始工作...")
}
func (c Camera) Stop()  {
	fmt.Println("相机 停止工作...")
}

//编写一个方法Working 方法，接收一个Usb接口类型变量
//只要是实现了 Usb接口 （所谓实现Usb接口，就是指实现了 Usb接口声明所有方法）
type Computer struct {

}

func (cp Computer) Working(usb Usb)  {
	usb.Start()
	usb.Stop()
}

func main()  {

	var computer Computer

	var phone Phone
	var camera Camera
	computer.Working(phone)
	computer.Working(camera)

}
