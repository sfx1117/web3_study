package main

import "fmt"

/*
	切片,声明和数组的区别是，数组是有固定长度的，切片没有长度
	1. 切片触发扩容前，切片一直共用相同的数组；
	2. 切片触发扩容后，会创建新的数组，并复制这些数据；
	3. 切片本身是一个特殊的指针，go 针对切片类型添加了一些语法糖，方便使用。
*/

func main() {
	//aa()
	//bb()
	//cc()
	//dd()
	//ee()
	ff()
}

// 声明切片
func aa() {
	//只声明，不初始化
	var s0 []int
	fmt.Println(s0)
	//1、声明并初始化一个空切片
	var s1 []int = []int{}
	fmt.Println(s1)
	//2、类型推导
	var s2 = []int{}
	fmt.Println(s2)
	//3、
	s3 := []int{1, 2, 3}
	fmt.Println(s3)
	//4、使用make函数 长度为0
	s4 := make([]int, 0)
	fmt.Println(s4)
	//5、使用make函数  长度为1，容量为3
	s5 := make([]int, 3, 5)
	fmt.Println(s5)
	//6、通过数组，初始化一个切片
	arr := [5]int{1, 2, 3, 4, 5}
	s6 := arr[1:]
	fmt.Println(s6)
	s7 := arr[:4]
	fmt.Println(s7)
	s8 := arr[2:3]
	fmt.Println(s8)
}

// 当基于同一个数组创建的切片，修改数组的值时，所有切片都受影响
func bb() {
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:]
	s2 := arr[1:3]
	s3 := arr[:3]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	//当修改arr的值时 s1,s2,s3都被修改
	arr[1] = 10
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

// 使用切片
func cc() {
	var s = []int{1, 2, 3, 4, 5}
	//使用下标访问切片
	e1 := s[0]
	e2 := s[1]
	fmt.Println(e1)
	fmt.Println(e2)

	//向指定位置赋值
	s[0] = 10
	s[1] = 20
	fmt.Println(s)

	//使用range 遍历切片
	for i, v := range s {
		fmt.Println(i, v)
	}

	//使用len()获取切片长度
	//ss := make([]int, 0, 5)
	//var ss = []int{}
	var ss []int
	l := len(ss)
	fmt.Println(l)
	//使用cap()获取切片容量
	c := cap(ss)
	fmt.Println(c)
}

// 使用append()
func dd() {
	var s []int
	fmt.Println(s)

	//使用append()向切片最后追加元素
	s = append(s)
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)

	//使用append向指定位置添加元素
	s = append(s[:1], append([]int{3}, s[1:]...)...)
	fmt.Println(s)

	//使用append删除指定位置元素
	s = append(s[:1], s[2:]...)
	fmt.Println(s)
}

// 使用copy()复制切片
func ee() {
	s1 := []int{1, 2, 3, 4}
	d1 := make([]int, 4, 5)
	fmt.Println(s1)
	fmt.Println(d1)

	//使用copy将s1复制到d1
	copy(d1, s1)
	fmt.Println(d1)

	s2 := []int{1, 2, 3, 4}
	d2 := make([]int, 3, 4)
	copy(d2, s2)
	fmt.Println(d2)
	fmt.Println(len(d2), cap(d2))
}

// append 扩容前
func ff() {
	var s = make([]int, 2, 3)
	s[0] = 1
	fmt.Println(s)

	s2 := append(s, 2)
	fmt.Println("s=", s)
	fmt.Println("s2=", s2, len(s2), cap(s2))

	s[1] = 2
	fmt.Println("s=", s)
	fmt.Println("s2=", s2, len(s2), cap(s2))

	appendFun(s)
	fmt.Println("s=", s)
	fmt.Println("s2=", s2, len(s2), cap(s2))

}
func appendFun(param []int) {
	param = append(param, 5)
	fmt.Println("param=", param)
	//param[]
}
