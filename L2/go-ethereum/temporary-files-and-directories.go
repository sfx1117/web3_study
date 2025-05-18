package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	f, err := os.CreateTemp("", "sample")
	check(err)

	fmt.Println("Temp file name:", f.Name())
	defer os.RemoveAll(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)
	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2, 3, 4}, 0644)
	check(err)

}

//Temp file name: C:\Users\sfx\AppData\Local\Temp\sample560200773
//Temp dir name: C:\Users\sfx\AppData\Local\Temp\sampledir420247762
