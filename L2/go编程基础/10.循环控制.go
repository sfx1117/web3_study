package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

/*
	循环控制
*/

func main() {
	//for1()
	//for2()
	//for3()
	//for4()
	//for5()
	//for6()
	//break1()
	//break2()
	//break3()
	//break4()
	break5()
}

// 方式1
func for1() {
	for i := 0; i < 10; i++ {
		fmt.Println("第", i+1, "次循环")
	}
}

// 方式2
func for2() {
	var i = 0
	for i < 10 {
		fmt.Println("第", i+1, "次循环")
		i++
	}
}

// 方式3 无限循环
func for3() {
	var ctx, _ = context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	var start bool
	var stop atomic.Bool
	for {
		if !start {
			start = true
			go func() {
				select {
				case <-ctx.Done():
					fmt.Println("ctx done")
					stop.Store(true)
					return
				}
			}()
		}
		fmt.Println("main")
		if stop.Load() {
			break
		}
	}
}

// 遍历数组
func for4() {
	var arr [10]int
	arr[0] = 0
	arr[1] = 1
	for i := range arr {
		fmt.Println(i)
	}
	for i, e := range arr {
		fmt.Println(i, e)
	}
}

// 遍历切片
func for5() {
	var s = []string{"for5", "test"}
	//s[0] = "for5"
	//s[1] = "test"
	var s1 = make([]string, 2)
	s1[0] = "test0"
	s1[1] = "test1"
	s1 = append(s1, "test2")
	for i := range s {
		fmt.Println(i)
	}

	for i, e := range s1 {
		fmt.Println(i, e)
	}
}

// 遍历map
func for6() {
	var m = make(map[int]string)
	m[0] = "test1"
	m[1] = "test2"
	m[2] = "test3"
	for i := range m {
		fmt.Println(i)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// break
func break1() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}

// 中断switch
func break2() {
	switch i := 1; i {
	case 1:
		fmt.Println(i)
		if i == 1 {
			break
		}
		fmt.Println("break")
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}
}

// 中断select
func break3() {
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("过了2秒")
	case <-time.After(time.Second):
		fmt.Println("过了1秒")
		break
		fmt.Println("brak之后")
	}
}

// for嵌套 不用leabl 跳出当前for
func break4() {
	for i := 0; i < 3; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			break
		}
	}
}

// for嵌套 使用leabl 跳到leabl处
func break5() {
	for {
	out:
		for i := 0; i < 3; i++ {
			fmt.Printf("不使用标记,外部循环, i = %d\n", i)
			for j := 0; j < 5; j++ {
				fmt.Printf("不使用标记,内部循环 j = %d\n", j)
				break out
			}
		}
	}
}
