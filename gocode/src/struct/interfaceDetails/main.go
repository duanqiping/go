package main

import "fmt"

type Student struct {

}

func (stu Student) Say()  {
	fmt.Println("stu say() ...")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =" ,i )
}

func (i integer) Hello() {
	fmt.Println("integer Say i =" ,i )
}

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Monster struct {

}
func (m Monster) Hello() {
	fmt.Println("Monster Hello()~~")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()~~")
}

func main()  {
	var stu Student
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b BInterface = i
	b.Hello()

	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster

	a2.Say()
	b2.Hello()
}

