package main

import "fmt"

/*
流程控制
*/

func main() {
	ifElse()
	sc()
}

// if语句
func ifElse() {
	var a int = 10
	if b := 1; a > 10 {
		b = 2
		fmt.Println("a > 10")
	} else if c := 3; b > 1 {
		b = 3
		fmt.Println("b > 1")
	} else {
		fmt.Println("other")
		if c == 3 {
			fmt.Println("c=3")
		}
		fmt.Println(a)
		fmt.Println(b)
	}
}

// switch case
func sc() {
	//1、基本用法
	var a string = "test string"
	switch a {
	case "aa":
		fmt.Println("a=aa")
	case "test":
		fmt.Println("a=test")
	case "dfd", "test string": // 可以匹配多个值，只要一个满足条件即可
		fmt.Println("a=dfd,或者a=test string")
	case "string":
		fmt.Println("a=string")
	default:
		fmt.Println("default case")
	}
	//2、变量b只在当前swithc模块中有效
	switch b := 5; b {
	case 1:
		fmt.Println("b=1")
	case 2, 3:
		fmt.Println("b=2 or b=3")
	case 4:
		fmt.Println("a=4")
	case 5:
		fmt.Println("a=5")

	}

	//3、不指定判断变量，直接在case中判断bool
	var b = 5
	switch {
	case a == "test":
		fmt.Println("a=test")
	case b == 3:
		fmt.Println("a=3")
	case b == 5, a == "test string":
		fmt.Println("b=5 or a=test string")
	}

	//4、仅适用于接口和泛型
	var d interface{}
	d = &a
	switch t := d.(type) {
	case byte:
		fmt.Println("d is byte type, ", t)
	case int:
		fmt.Println("d is int type, ", t)
	case *byte:
		fmt.Println("d is *byte type, ", t)
	case *int:
		fmt.Println("d is *int type, ", t)
	case string:
		fmt.Println("d is string type, ", t)
	case *string:
		fmt.Println("d is *string type, ", t)
	default:
		fmt.Println("d is unknown type, ", t)
	}
}
