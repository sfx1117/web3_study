package main

import "fmt"

/*
	局部变量:在函数内定义的变量
*/

func main() {
	var s string = "hello world"
	var a int = 10
	var b bool = true

	//1、类型推导
	s1 := "类型推导"
	//2、完整的变量声明方式
	var a1 int = 100
	//3、只声明，不初始化
	var b1 bool

	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s1)
	fmt.Println(a1)
	fmt.Println(b1)

	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
}

//4、在函数的返回值值中声明
func test1() (a int, b string) {
	return 10, "test1"
}

func test2() (a int, b string) {
	a = 1
	b = "test2"
	return
}
func test3() (a int, b string) {
	return
}
