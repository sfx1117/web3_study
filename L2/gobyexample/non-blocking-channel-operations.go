package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan bool)
	select {
	case msg := <-c1:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	select {
	case c1 <- "msg":
		fmt.Println("sent message", "msg")
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-c1:
		fmt.Println("received message", msg)
	case sig := <-c2:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

//no message received
//no message sent
//no activity
