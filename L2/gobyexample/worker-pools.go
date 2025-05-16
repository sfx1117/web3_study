package main

import (
	"fmt"
	"time"
)

func work(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second * 2)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const n = 5
	jobs := make(chan int, n)
	results := make(chan int, n)

	for i := 1; i < 4; i++ {
		go work(i, jobs, results)
	}
	for j := 1; j <= n; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= n; a++ {
		fmt.Println(<-results)
	}
}

//worker 2 started  job 2
//worker 1 started  job 1
//worker 3 started  job 3
//worker 1 finished job 1
//worker 1 started  job 4
//worker 3 finished job 3
//worker 2 finished job 2
//2
//worker 3 started  job 5
//6
//4
//worker 1 finished job 4
//worker 3 finished job 5
//8
//10
