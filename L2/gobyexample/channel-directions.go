package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(ping <-chan string, po chan<- string) {
	msg := <-ping
	po <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	fmt.Println(<-pings)
	ping(pings, "passed message2")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

//passed message
//passed message2
