package main    // 声明 main 包

import (
    "fmt"       // 导入 fmt 包，打印字符串是需要用到
	"math"
)
func GetData() (int, int) {
	return 100, 200
}

func main() {
	// 声明 main 主函数
	var str = "C语言中文网\nGo语言教程"
	fmt.Println(str)

    fmt.Println("Hello World!") // 打印 Hello World!

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	
	// 声明 hp 变量
	// var hp = 10
	// 再次声明并赋值
	//a, _ := GetData()
	//_, b := GetData()
	//fmt.Println(a, b)

	//var a int = 100
	//var b int = 200
	//b, a = a, b
	//fmt.Println(a, b)

	//声明局部变量 a 和 b 并赋值
	var a int = 3
	var b int = 4
	//声明局部变量 c 并计算 a 和 b 的和
	c := a + b
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17
	var damage = float32(attack-defence) * damageRate
	fmt.Println(damage)
}