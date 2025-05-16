package main

import "fmt"

func sum(nums ...int) {
	fmt.Println("nums:", nums)
	var total int
	for _, n := range nums {
		total += n
	}
	fmt.Println("total=", total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}

//nums: [1 2]
//total= 3
//nums: [1 2 3]
//total= 6
//nums: [1 2 3 4 5]
//total= 15
