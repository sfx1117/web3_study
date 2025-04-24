package main

import "fmt"

//基本数据类型
func main() {
	// 1. 整型
	// 有符号：int,int8,int16,int32,int64
	//无符号：uint,uint8,uint16,uint32,uint64
	//特殊：uintptr

	var a int8 = 10
	var b uint8 = 10
	fmt.Println(a, b)

	//2、浮点数
	// float32 有效数字7位,float64 有效数字15位
	var c float32 = 1.1234561
	var d float64 = 1.12345678
	fmt.Println(c, d)

	//3、布尔值
	var e bool = true
	fmt.Println(e)

	//4、byte  是uint8的内置别名，字符串可以直接转为[]byte
	var s string = "hello word"
	var bs []byte = []byte(s)
	fmt.Println(bs)

	var bs1 []byte = []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 100}
	var s1 string = string(bs1)
	fmt.Println(s1)
	fmt.Println(s1 == s)

	//5、rune类型 是int32的内置别名
	var r1 rune = '中'
	var r2 rune = 'a'
	fmt.Println(r1, r2)

	var s3 string = "中国--广州"
	var r3 []rune = []rune(s3)
	fmt.Println(r3)

	//6、字符串
	var str string = "hello word"
	fmt.Println(str)

	//7、string byte rune之间的相互转换
	var s4 string = "GO语言"
	var bs4 []byte = []byte(s4)
	var r4 []rune = []rune(s4)
	fmt.Println(s4, "length=", len(s4))
	fmt.Println(bs4, "length=", len(bs4))
	fmt.Println(r4, "length=", len(r4))

}
