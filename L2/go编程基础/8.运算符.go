package main

import (
	"fmt"
)

/*
运算符
*/

func main() {
	//1、加减乘除 取模
	var a, b = 1, 2
	var sum = a + b
	var sub = a - b
	var mul = a * b
	var div = a / b
	var mod = a % b
	fmt.Println(sum, sub, mul, div, mod)
	//自增自减
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
	//错误写法，不可以用自增自减时再做运算，也不可以用自增自减赋值
	//++a
	//--a
	//var c=a++
	//var c=a++ +1
	//不同类型混合计算时，需要先转换类型
	var c = 10 + 0.1
	var d byte = 1
	fmt.Println(c, d)
	var e = c + float64(d)
	fmt.Println(e)

	//2、关系运算符
	var aa int = 1
	var bb int = 2
	fmt.Println(aa == bb, aa != bb, aa > bb, aa < bb, aa >= bb, aa <= bb)

	//3、逻辑运算符
	var aaa bool = true
	var bbb bool = false
	fmt.Println(aaa && bbb, aaa || bbb, !(aaa && bbb))

	//4、位运算符
	fmt.Println(0&1, 0|1, 0^1)
	fmt.Println(0&0, 0|0, 0^0)
	fmt.Println(1&1, 1|1, 1^1)
	fmt.Println(1&0, 1|0, 1^0)

	//5、赋值运算符
	var a5, b5 = 1, 2
	var c5 int
	c5 = a5 + b5
	fmt.Println(c5)
	c5 += a5
	fmt.Println(c5)
	c5 -= a5
	fmt.Println(c5)
	c5 *= a5
	fmt.Println(c5)
	c5 /= a5
	fmt.Println(c5)
	c5 %= a5
	fmt.Println(c5)
	c5 <<= a5 //左移1位
	fmt.Println(c5)
	c5 >>= a5 //右移一位
	fmt.Println(c5)
	c5 &= a5
	fmt.Println(c5)
	c5 |= a5
	fmt.Println(c5)
	c5 ^= a5
	fmt.Println(c5)

	//6、其他运算符
	var a6 int = 1
	var pa *int
	fmt.Println(a6)
	fmt.Println(pa)
	pa = &a6 //p6=a6的指针
	fmt.Println(pa)
	fmt.Println(*pa)

	//7、运算优先级
	var a7, b7, c7, d7 = 1, 2, 3, 4
	var e7 int
	e7 = (a7 + b7) * c7 / d7
	fmt.Println("e7=", e7)
	e7 = (a7 + b7) * (c7 / d7)
	fmt.Println("e7=", e7)
	//先算3 + 4 ^ 3
	// 3+4=7,7^3=4
	//再算2&2*3<<1
	//2&2=2,2*3=6,6<<1=12
	//最后再算4|12=12
	f := 3 + 4 ^ 3 | 2&2*3<<1
	fmt.Println(f == 12)
}
