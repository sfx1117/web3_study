package main

import "fmt"

func zeroval(i int) {
	i = 0
}
func zeroport(i *int) {
	*i = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroport(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}

//initial: 1
//zeroval: 1
//zeroptr: 0
//pointer: 0xc00000a0f8
