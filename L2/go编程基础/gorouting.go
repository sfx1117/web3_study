package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	gorouting 和 channel
*/

func main() {
	//aa()
	bb()
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(s)
	}
}
func aa() {
	//方式一 无参
	go func() {
		fmt.Println("run goroutine in closure")
	}()
	//方式二  有参数
	go func(s string) {
		fmt.Println(s)
	}("gorouine: closure params")
	//方式三
	go say("in goroutine: world")
	say("hello")
}

//gorouting的并发问题
//线程安全计数结构体
type Safe struct {
	count int
	mu    sync.Mutex
}

func (s *Safe) add() {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.count++
}
func (s *Safe) getCount() int {
	defer s.mu.Unlock()
	s.mu.Lock()
	return s.count
}

//线程不安全计数结构体
type Unsafe struct {
	count int
}

func (u *Unsafe) add() {
	u.count++
}
func (u *Unsafe) getCount() int {
	return u.count
}
func bb() {
	//u := Unsafe{}
	s := Safe{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				//u.add()
				s.add()
			}
		}()
	}

	time.Sleep(time.Second)
	//fmt.Println(u.getCount())
	fmt.Println(s.getCount())
}
