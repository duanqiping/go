package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{})  {
	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	//2. 获取到 reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue vlaue 值为：%v,rValue type 为：%T\n",rValue,rValue)

	n2 := rValue.Int() + 10
	//n2 := rValue.Float() + 10//panic: reflect: call of reflect.Value.Float on int Value
	fmt.Println(n2)

	//下面我们将 rVal 转成 interface{}
	Iv := rValue.Interface()
	num2,ok := Iv.(int)
	if ok {
		fmt.Println(num2,ok)
	}else{
		fmt.Println("传入的参数不是整形")
	}
}

func reflect02(b interface{})  {
	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//3. 获取 变量对应的Kind
	//(1) rVal.Kind() ==>
	kind1 := rVal.Kind()
	kind2 := rTyp.Kind()
	fmt.Printf("kind =%v kind=%v\n", kind1, kind2)

	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言.
	//同学们可以使用 swtich 的断言形式来做的更加的灵活

	stu,ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

type Student struct {
	Name string
	Age int
}

type Monster struct {
	Name string
	Age int
}

func main()  {
	//var num  = 100
	//reflect01(num)

	stu := Student{
		Name:"孙悟空",
		Age:500,
	}
	reflect02(stu)

}
