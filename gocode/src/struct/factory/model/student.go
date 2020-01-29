package model

type student struct {
	Name string
	age int
}


//因为student结构体首字母是小写，因此是只能在model使用
//我们通过工厂模式来解决
func GetNewStudent(name string,age int) *student {
	return &student{name,age}
}

//返回年龄
func (s *student) GetNewAge() int {
	return s.age
}