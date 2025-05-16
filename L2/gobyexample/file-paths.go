package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("aaa", "bbb", "ccc")
	fmt.Println(p)
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	fmt.Println(filepath.IsAbs("C:/文档"))

	fileName := "config.json"
	ext := filepath.Ext(fileName)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(fileName, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}

//aaa\bbb\ccc
//dir1\filename
//dir1\filename
//Dir(p): aaa\bbb
//Base(p): ccc
//false
//false
