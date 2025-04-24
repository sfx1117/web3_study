package main

import "fmt"

/*
	全局变量:在函数外定义的变量称为全局变量
*/

var s string = "hello"
var a int = 10
var b bool = true

var (
	s1 string = "word"
	a1 int    = 100
	b1 bool   = false
)

var s2 string
var a2 int
var b2 bool

func main() {
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s1)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(s2)
	fmt.Println(a2)
	fmt.Println(b2)

}
