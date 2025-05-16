package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()
	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(time.Second):
		fmt.Println("time out1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 3):
		fmt.Println("tiemout2")
	}
}

//time out1
//two
