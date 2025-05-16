package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	var sum int
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum:", sum)

	for i, n := range nums {
		if n == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{
		"a": "a",
		"b": "b",
	}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

//sum: 6
//index: 2
//a -> a
//b -> b
//key: a
//key: b
//0 103
//1 111
