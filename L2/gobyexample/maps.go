package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	v3 := m["k3"]
	fmt.Println("v3:", v3)
	fmt.Println("len:", len(m))

	delete(m, "k1")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{
		"floo": 1,
		"dfd":  3,
	}

	n2 := map[string]int{
		"floo": 1,
		"dfd":  3,
	}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

//map: map[k1:1 k2:2]
//v1: 1
//v3: 0
//len: 2
//map: map[k2:2]
//map: map[]
//prs: false
//n == n2
