package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		upper := strings.ToUpper(scanner.Text())
		fmt.Println(upper)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

}

//dfdsg
//DFDSG
//dfsre
//DFSRE
//3231bd
//3231BD
