package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Println(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	file := createFile("C:/文档/web3/L2/gobyexample/defer.txt")
	defer closeFile(file)
	writeFile(file)

}

//creating
//writing
//&{0xc00006e6c8} data
//closing
