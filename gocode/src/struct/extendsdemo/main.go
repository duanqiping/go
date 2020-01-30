package main

import "fmt"

type Student struct {
	Name string
	Age byte
	Score float32
}

//将Pupil 和 Graduate 共有的方法也绑定到 *Student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score float32) {
	//业务判断
	stu.Score = score
}

//给 *Student 增加一个方法，那么 Pupil 和 Graduate都可以使用该方法
func (stu *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

type Pupil struct {
	Student //嵌入了Student匿名结构体
}

//这时Pupil结构体特有的方法，保留
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中.....")
}

//大学生
type Graduate struct {
	Student //嵌入了Student匿名结构体
}

//显示他的成绩
//这时Graduate结构体特有的方法，保留
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中.....")
}

func main()  {
	pupil := Pupil{}
	pupil.Name = "qiping"
	pupil.Age = 29
	pupil.Student.Score = 88
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
	fmt.Println("res=", pupil.Student.GetSum(1, 2))
	fmt.Println(pupil)

	graduate := &Graduate{}
	graduate.Student.Name = "mary~"
	graduate.Student.Age = 28
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
	fmt.Println("res=", graduate.Student.GetSum(10, 20))
}
