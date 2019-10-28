
package main    // 声明 main 包

import (
    "fmt"       // 导入 fmt 包，打印字符串是需要用到
    "math"
)

type Weekday int

const Pi64 float64 = math.Pi

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main() {   // 声明 main 主函数
    fmt.Println("Hello World!") // 打印 Hello World!
    fmt.Println(Sunday) // 打印 Hello World!
    fmt.Println(Monday) // 打印 Hello World!
    fmt.Println(Tuesday) // 打印 Hello World!

    var x float32 = float32(Pi64)
    var y float64 = Pi64
    var z complex128 = complex128(Pi64)
    fmt.Println(x) // 打印 Hello World!
    fmt.Println(y) // 打印 Hello World!
    fmt.Println(z) // 打印 Hello World!
}

