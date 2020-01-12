package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	//统计字符串的长度，按字节 len(str)
	////golang的编码统一为utf-8 (ascii的字符(字母和数字) 占一个字节，汉字占用3个字节)

	str := "hello北京"
	fmt.Println("str的长度为：",len(str))

	//字符串遍历，同时处理有中文的问题 r := []rune(str)
	str2 := []rune(str)
	for i := 0; i<len(str2); i++  {
		fmt.Printf("%c\n",str2[i])
	}

	//for _,val := range str{
	//	fmt.Printf("%c\n",val)
	//}

	//字符串转整数:	 n, err := strconv.Atoi("12")
	n,err :=strconv.Atoi("123")
	if err != nil {
		fmt.Println("错误的原因是",err)
	}else{
		fmt.Println("转换后的结果是",n)
	}

	//4)整数转字符串  str = strconv.Itoa(12345)
	str3 := strconv.Itoa(8976)
	fmt.Printf("str的值为%v 类型为%T\n",str3,str3)

	//5)字符串 转 []byte:  var bytes = []byte("hello go")
	var bytes = []byte("hello go 北京")
	fmt.Printf("str=%v\n",bytes)

	//6)[]byte 转 字符串: str = string([]byte{97, 98, 99})
	str = string([]byte{99,98,97})
	fmt.Printf("str=%v\n", str)

	//10进制转 2, 8, 16进制:  str = strconv.FormatInt(123, 2),返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是=%v\n", str)

	//查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
	b := strings.Contains("seafood", "mary")
	fmt.Printf("b=%v\n", b)

	//统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "e") //4
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)

	//10)不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true

	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b) //true

	//11)返回子串在字符串第一次出现的index值，如果没有返回-1 :
	//strings.Index("NLT_abc", "abc") // 4

	index := strings.Index("NLT_abcabcabc", "abcg") // 4
	fmt.Printf("index=%v\n",index)

	//12)返回子串在字符串最后一次出现的index，
	//如没有返回-1 : strings.LastIndex("go golang", "go")

	index = strings.LastIndex("go golang", "go") //3
	fmt.Printf("index=%v\n",index)

	//将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go语言", n)
	//n可以指定你希望替换几个，如果n=-1表示全部替换

	str5 := "go go hello"
	str = strings.Replace(str5, "go", "北京", -1)
	fmt.Printf("str=%v str5=%v\n", str, str)

	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：
	//strings.Split("hello,wrold,ok", ",")
	strArr := strings.Split("hello,wrold,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	//15)将字符串的字母进行大小写的转换:
	//strings.ToLower("Go") // go strings.ToUpper("Go") // GO

	str = "goLang Hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str) //golang hello

	//将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gopher ntrn   ")
	str = strings.TrimSpace(" tn a lone gopher ntrn   ")
	fmt.Printf("str=%q\n", str)

	//17)将字符串左右两边指定的字符去掉 ：
	//strings.Trim("! hello! ", " !")  // ["hello"] //将左右两边 ! 和 " "去掉
	str = strings.Trim("! he!llo! ", " !")
	fmt.Printf("str=%q\n", str)

	//20)判断字符串是否以指定的字符串开头:
	//strings.HasPrefix("ftp://192.168.10.1", "ftp") // true

	b = strings.HasPrefix("ftp://192.168.10.1", "hsp") //true
	fmt.Printf("b=%v\n", b)

}
