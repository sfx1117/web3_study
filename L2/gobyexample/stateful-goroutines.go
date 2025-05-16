package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readop struct {
	key  int
	resp chan int
}
type writeop struct {
	key   int
	value int
	resp  chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readop)
	writes := make(chan writeop)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case r := <-reads:
				r.resp <- state[r.key]
			case w := <-writes:
				state[w.key] = w.value
				w.resp <- true
			}
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := readop{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeop{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

//readOps: 62124
//writeOps: 6218
