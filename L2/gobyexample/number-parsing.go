package main

import (
	"fmt"
	"strconv"
)

func main() {
	f1, _ := strconv.ParseFloat("1.2324", 64)
	fmt.Println(f1)

	i, _ := strconv.ParseInt("232", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}

//1.2324
//232
//456
//789
//135
//strconv.Atoi: parsing "wat": invalid syntax
