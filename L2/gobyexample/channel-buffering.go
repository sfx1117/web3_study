package main

import "fmt"

func main() {
	msg := make(chan string, 2)
	msg <- "channel"
	msg <- "channel2"
	//msg <- "channel3"
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}

//channel
//channel2
