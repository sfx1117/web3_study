package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	fmt.Println(fact(7))

	var fib func(a int) int

	fib = func(a int) int {
		if a < 2 {
			return a
		}
		return fib(a-1) + fib(a-2)
	}
	fmt.Println(fib(7))
}

//5040
//13
