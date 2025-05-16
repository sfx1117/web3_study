package main

import "fmt"

func val() (int, int) {
	return 3, 7
}

func main() {
	i, i2 := val()
	fmt.Println(i)
	fmt.Println(i2)

	_, i4 := val()
	fmt.Println(i4)
}

//3
//7
//7
