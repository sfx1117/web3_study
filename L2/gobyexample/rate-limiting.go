package main

import (
	"fmt"
	"time"
)

func main() {
	req := make(chan int, 5)
	for i := 0; i < 5; i++ {
		req <- i
	}
	close(req)

	ticker := time.NewTicker(time.Millisecond * 200)
	for re := range req {
		<-ticker.C
		fmt.Println("request", re, time.Now())
	}

	buLim := make(chan time.Time, 3)
	for j := 0; j < 3; j++ {
		buLim <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			buLim <- t
		}
	}()

	busReq := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		busReq <- i
	}
	close(busReq)
	for bus := range busReq {
		<-buLim
		fmt.Println("request", bus, time.Now())
	}
}

//request 0 2025-05-14 16:05:17.8044768 +0800 CST m=+0.200714401
//request 1 2025-05-14 16:05:18.0040128 +0800 CST m=+0.400250401
//request 2 2025-05-14 16:05:18.2045613 +0800 CST m=+0.600798901
//request 3 2025-05-14 16:05:18.4040157 +0800 CST m=+0.800253301
//request 4 2025-05-14 16:05:18.6040053 +0800 CST m=+1.000242901
//request 1 2025-05-14 16:05:18.6040053 +0800 CST m=+1.000242901
//request 2 2025-05-14 16:05:18.6040053 +0800 CST m=+1.000242901
//request 3 2025-05-14 16:05:18.6040053 +0800 CST m=+1.000242901
//request 4 2025-05-14 16:05:18.8048654 +0800 CST m=+1.201103001
//request 5 2025-05-14 16:05:19.0046617 +0800 CST m=+1.400899301
