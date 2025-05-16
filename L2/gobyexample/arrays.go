package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)
	a[3] = 3
	fmt.Println("set:", a)
	fmt.Println("get:", a[3])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 3, 4: 104}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d: ", twoD)
}

//emp: [0 0 0 0 0]
//set: [0 0 0 3 0]
//get: 3
//len: 5
//dcl: [1 2 3 4 5]
//dcl: [1 2 3 4 5]
//dcl: [100 0 0 3 104]
//2d:  [[0 1 2] [1 2 3]]
//2d:  [[1 2 3] [4 5 6]]
