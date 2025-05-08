package main

import (
	"fmt"
	"strconv"
)

/*
	类型转换
	byte=uint8
	当两个结构体中的字段名称以及类型都完全相同，仅结构体名称不同时，这两个结构体类型即可相互转换。
*/

func main() {
	//aa()
	//bb()
	//cc()
	//dd()
	ee()
}

//数字类型转换
func aa() {
	var a int = 7
	var b byte = 3
	var c float32

	c = float32(a) / float32(b)
	fmt.Println(c)

	//当int32强转成byte时，超过位数会被截断
	var e int32 = 255
	var f byte = byte(e)
	fmt.Println(e)
	fmt.Println(f)
}

//字符串类型转换
func bb() {
	//string 转换为[]byte、[]rune
	var str string = "hello 你好！"
	var bt []byte = []byte(str)
	var ru []rune = []rune(str)
	fmt.Println(str)
	fmt.Println(bt)
	fmt.Println(ru)
	//[]byte、[]rune转换为string
	str1 := string(bt)
	str2 := string(ru)
	fmt.Println(str1)
	fmt.Println(str2)

	//string 转换为数字
	ss := "123"
	num, err := strconv.Atoi(ss)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)

	//将数字转换为string
	sss := strconv.Itoa(num)
	fmt.Println(sss)
}

//interface类型转换
func cc() {
	var i interface{} = "3"
	num, ok := i.(string)
	if ok {
		fmt.Println(num)
	} else {
		fmt.Println("conversion failed")
	}

	//使用switch case
	switch t := i.(type) {
	case int:
		fmt.Println("int", t)
	case string:
		fmt.Println("string", t)
	case byte:
		fmt.Println("byte", t)
	}

}

//接口类型和结构体类型之间互转
type Interf interface {
	Get() string
}
type Stru struct {
	name string
}

func (s *Stru) Get() string {
	return s.name
}
func dd() {
	//结构体实现了接口 通过结构体声明接口类型
	var i Interf = &Stru{"sfx"}
	fmt.Println(i)

	// 接口类型  转换为结构体
	s, err := i.(*Stru)
	fmt.Println(s, err)
}

//结构体之间的转换
type AAA struct {
	name string
	age  int
}
type BBB struct {
	name string
	age  int
}

func (b *BBB) getName() string {
	return b.name
}
func ee() {
	//声明AAA
	a := AAA{
		name: "sfx",
		age:  11,
	}
	//将AAA转换为BBB
	b := BBB(a)
	fmt.Println(b.getName())

	//只能结构体类型实例之间相互转换，指针不可以相互转换
	//如果强转的话，程序不报错，但是结果是转换失败的
	var c interface{} = &a
	d, ok := c.(*BBB)
	fmt.Println(d, ok)

}
