package main

import (
	"fmt"
	"unsafe"
)

/*
	指针
*/

var s string = "string"
var ps *string = &s

var a int = 10
var pa *int = &a

var b bool
var pb *bool
var pb1 *bool = &b

func main() {
	fmt.Println(s, ps, *ps)
	// fmt.Println(ps)
	fmt.Println(a, pa, *pa)
	*pa = 100
	fmt.Println(a, pa, *pa)
	// fmt.Println(pa)
	fmt.Println(b, pb, pb1, *pb1)
	// fmt.Println(pb)
	// fmt.Println(pb1)

	//修改指针的值
	var pps **string = &ps
	fmt.Println(pps)

	// 指针、unsafe.Pointer 和 uintptr
	//指针不能参与直接计算，否则直接报错
	var aa int = 10
	var paa *int = &aa
	// paa = paa + 1 //mismatched types *int and untyped int
	fmt.Println(paa)

	//把指针转换成unsafe.Point
	var up1 = unsafe.Pointer(paa)
	fmt.Println(up1)
	//把unsafe.Pointer转换成uintptr
	var uptr uintptr = uintptr(up1)
	fmt.Println(uptr)

}
