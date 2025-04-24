package main

import "fmt"

/*
	定义多变量
*/

var s, a, b = "hello", 1, true
var c, d int = 1, 2
var e, f bool

func main() {
	fmt.Println(s, a, b, c, d, e, f)
	var g, h = 12, "string"
	var i, j int
	l, m, n := 13, true, "string"
	fmt.Println(g, h, i, j, l, m, n)
}
