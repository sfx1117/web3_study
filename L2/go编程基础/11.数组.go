package main

import "fmt"

/*
	数组
*/

func main() {
	//aa()
	//bb()
	//cc()
	//作为参数
	a := [5]int{1, 2, 3, 4, 5}
	//receiveArr(a)
	receiveArrPoint(&a)
	fmt.Println(a)
}

// 声明数组
func aa() {
	//方式一  仅声明
	var a [2]int
	var aa [2]map[string]int
	fmt.Println(a)
	fmt.Println(aa)

	//方式二  声明并初始化
	var b [2]int = [2]int{1, 2}
	var bb [2]map[string]int = [2]map[string]int{
		{"a": 1},
		{"b": 2},
	}
	fmt.Println(b)
	fmt.Println(bb)

	//方式三 类型推导
	var c = [3]int{1, 2, 3}
	cc := [3]string{"1", "2", "3"}
	fmt.Println(c)
	fmt.Println(cc)

	//方式四  使用...代替数组长度
	d := [...]int{1, 2}
	fmt.Println(d, len(d))

	//方式五 声明时初始化指定下标的值
	e := [5]string{1: "aaa", 2: "bbb"}
	fmt.Println(e, len(e))
}

// 访问数组
func bb() {
	a := [5]int{1, 2, 3, 4, 5}
	//1、使用下标获取元素
	b := a[1]
	fmt.Println(b)

	//2、使用range 遍历数组
	for i, v := range a {
		fmt.Println(i, a[i], v)
	}
	for i := range a {
		fmt.Println(i, a[i])
	}

	//3、len()获取数组长度
	c := len(a)
	fmt.Println(c)

	for j := 0; j < len(a); j++ {
		fmt.Println(j)
	}
}

// 多维数组
func cc() {
	//1、2维数组
	aa := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(aa)

	//3维数组
	bbb := [2][3][2]int{
		{
			{1, 2},
			{3, 4},
			{5, 6},
		},
		{
			{7, 8},
			{9, 10},
			{11, 12},
		},
	}
	fmt.Println(bbb)

	//也可以先声明，后面根据下标赋值
	var ccc = [2][2][2]int{}
	ccc[0][0][0] = 111
	ccc[1][1][1] = 222
	fmt.Println(ccc)

	//多维数组遍历是需要嵌套遍历
	for i, v := range bbb {
		fmt.Println(i, v)
		for j, jv := range v {
			fmt.Println(j, jv)
			for n, nv := range jv {
				fmt.Println(n, nv)
			}
		}
	}

}

// 数组作为参数，跟基本类型一样，只有将数组的指针作为参数，才能修改外部实参的值
func receiveArr(param [5]int) {
	param[1] = 10
	fmt.Println(param)
}
func receiveArrPoint(param *[5]int) {
	param[1] = 10
	fmt.Println(param)
}
