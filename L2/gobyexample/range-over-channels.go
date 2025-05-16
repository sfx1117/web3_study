package main

import "fmt"

func main() {
	c1 := make(chan string, 2)
	c1 <- "111"
	c1 <- "222"
	close(c1)
	for e := range c1 {
		fmt.Println(e)
	}
}

//111
//222
