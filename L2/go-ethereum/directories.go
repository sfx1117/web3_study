package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	err := os.Mkdir("subdir", 0755)
	check(err)
	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	createEmptyFile("subdir/file1")
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check(err)
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)
	c, err = os.ReadDir(".")
	check(err)
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}

//Listing subdir/parent
//child true
//file2 false
//file3 false
//Listing subdir/parent/child
//file4 false
//Visiting subdir
//subdir true
//subdir\file1 false
//subdir\parent true
//subdir\parent\child true
//subdir\parent\child\file4 false
//subdir\parent\file2 false
//subdir\parent\file3 false
