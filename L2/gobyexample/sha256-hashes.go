package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"
	hash := sha256.New()
	hash.Write([]byte(s))
	bs := hash.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

//sha256 this string
//1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb
