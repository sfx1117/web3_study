package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range,", i)
	}

	for {
		fmt.Println("loop")
		break
	}
	for n := range 6 {
		if n%2 == 0 {
			fmt.Println(n)
			continue
		}
	}

}

//0
//1
//2
//0
//1
//2
//range, 0
//range, 1
//range, 2
//loop
//0
//2
//4
