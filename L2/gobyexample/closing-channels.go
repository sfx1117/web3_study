package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 5)
	c2 := make(chan bool)

	go func() {
		for {
			if j, ok := <-c1; ok {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				c2 <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		c1 <- i
		fmt.Println("sent job", i)
	}
	close(c1)
	fmt.Println("sent all jobs")

	fmt.Println("<-c2:", <-c2)

	_, ok := <-c1
	fmt.Println("received more jobs:", ok)
}

//sent job 0
//sent job 1
//sent job 2
//sent all jobs
//received job 0
//received job 1
//received job 2
//received all jobs
//<-c2: true
//received more jobs: false
