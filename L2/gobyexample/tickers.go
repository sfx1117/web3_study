package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	c1 := make(chan bool)

	go func() {
		for {
			select {
			case <-c1:
				fmt.Println("c1 receive true")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(time.Second * 3)
	ticker.Stop()
	c1 <- true
	fmt.Println("Ticker stopped")

}

//Tick at 2025-05-14 15:23:15.076885 +0800 CST m=+1.000000001
//Tick at 2025-05-14 15:23:16.076885 +0800 CST m=+2.000000001
//Tick at 2025-05-14 15:23:17.076885 +0800 CST m=+3.000000001
//c1 receive true
//Ticker stopped
