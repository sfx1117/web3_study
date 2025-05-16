package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"a", "d", "b", "c"}
	fmt.Println("排序前", strs)
	slices.Sort(strs)
	fmt.Println("排序后", strs)

	ints := []int{3, 7, 5, 11}
	fmt.Println("排序前", ints)
	slices.Sort(ints)
	fmt.Println("排序后", ints)

	isSorted := slices.IsSorted(ints)
	fmt.Println("是否已排序", isSorted)
}

//排序前 [a d b c]
//排序后 [a b c d]
//排序前 [3 7 5 11]
//排序后 [3 5 7 11]
//是否已排序 true
