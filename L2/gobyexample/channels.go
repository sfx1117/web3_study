package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "channel"
	}()

	msg := <-messages
	fmt.Println(msg)

}

//channel
