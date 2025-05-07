package main

import (
	"context"
	"fmt"
	"time"
)

/*
	select 关键字
*/

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
	}()

	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	go func() {
		for {
			select {
			case x := <-ch1:
				fmt.Println("ch1----", x)
			case y := <-ch2:
				fmt.Println("ch2----", y)
			case z := <-ch3:
				fmt.Println("ch3----", z)
			case <-ctx.Done():
				fmt.Print("timeout")
				if err := ctx.Err(); err != nil {
					//todo
				}
				return
			}
		}
	}()
	time.Sleep(4 * time.Second)
}
