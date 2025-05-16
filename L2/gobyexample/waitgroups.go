package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(i)
		}()
	}

	wg.Wait()
}

//Worker 2 starting
//Worker 1 starting
//Worker 0 starting
//Worker 0 done
//Worker 2 done
//Worker 1 done
