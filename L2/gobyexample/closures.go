package main

import "fmt"

//闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	next := intSeq()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	next1 := intSeq()
	fmt.Println(next1())
}

//1
//2
//3
//1
