package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	strs := []string{"peach", "banana", "kiwi"}
	fmt.Println("排序前", strs)
	lenCmp := func(a string, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(strs, lenCmp)
	fmt.Println(strs)

	type person struct {
		name string
		age  int
	}

	pp := []person{
		{name: "111", age: 23},
		{name: "222", age: 35},
		{name: "333", age: 31},
	}
	slices.SortFunc(pp, func(a person, b person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(pp)

}

//排序前 [peach banana kiwi]
//[kiwi peach banana]
//[{111 23} {333 31} {222 35}]
