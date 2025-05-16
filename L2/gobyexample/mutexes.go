package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu   sync.Mutex
	coun map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.coun[name]++
}

func main() {
	c := Container{
		coun: map[string]int{
			"a": 0,
			"b": 0,
		},
	}
	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			c.inc(name)
		}
	}
	wg.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)
	wg.Wait()
	fmt.Println(c.coun)

}

//map[a:2000 b:1000]
