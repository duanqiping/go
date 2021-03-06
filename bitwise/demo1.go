package main

import "fmt"

//位运算
func main()  {

	//计算机的运算实际上 就是 补码的运算
	//正数 的 反码和补码 和 原吗一致
	//负数-》符号位不变，其他位取反（反码）-》反码加1（补码）

	fmt.Println(2&3)  // 0000 0010  0000 0011=> 0000 0010 => 2
	fmt.Println(2|3)  // 0000 0010  0000 0011=> 0000 0011 => 3
	fmt.Println(2^3)  // 0000 0010  0000 0011=> 0000 0001 => 1
	fmt.Println(-2|3) // 1000 0010  0000 0011=> 1111 1101  0000 0011=》1111 1110  0000 0011 =》1111 1111（补码）-》1111 1110-》1000 0001=》 -1
	fmt.Println(-2&3) // 1000 0010  0000 0011=> 1111 1101  0000 0011=》1111 1110  0000 0011 =》0000 0010（补码）-》2

	fmt.Println( 1>>2) // 0000 0001 => 0000 0000 => 0
	fmt.Println(1<<2)  // 0000 0001 => 0000 0100 => 4
	fmt.Println(-1<<2) // 1000 0001 => 1000 0100 =>-4
}
