package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000
	const d = 3e20 / n
	fmt.Println(3e20)
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))
}

//constant
//3e+20
//6e+14
//600000000000000
//0.9538537219686818
