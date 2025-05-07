package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	map
	当并发读写map时，需要添加互斥锁，否则会报错
*/

func main() {
	//aa()
	//bb()
	//cc()
	//dd()
	ee()
	//ff()
}

// 声明map
func aa() {
	//只声明，不初始化
	var m1 map[string]string
	fmt.Println(m1)

	m2 := map[string]string{}
	fmt.Println(m2)

	m3 := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	fmt.Println(m3)

	m4 := make(map[string]int)
	fmt.Println(m4)

	m5 := make(map[string]int, 3)
	fmt.Println(m5)

}

// 使用map
func bb() {
	var m = make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5

	//根据k获取v
	v := m["1"]
	fmt.Println(v)
	v1, ex := m["1"]
	fmt.Println(v1, ex)
	v2, ex2 := m["10"]
	fmt.Println(v2, ex2)

	//修改值
	m["2"] = 20
	fmt.Println(m)

	//获取map的长度  （多少个键值对）
	count := len(m)
	fmt.Println(count)

	//使用delete函数，删除元素
	delete(m, "10")

	//遍历map
	for k, v := range m {
		fmt.Println(k, v)
	}

	//在遍历map时 delete
	for k := range m {
		delete(m, k)
	}
	fmt.Println(m)
}

// 当map作为入参,会更改原始数据
func cc() {

	var m = make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	fmt.Println(m)
	receiveMap(m)
	fmt.Println(m)
}

func receiveMap(m map[string]int) {
	m["3"] = 3
	m["2"] = 20
}

// 并发使用map ，不加锁
// 报错 concurrent map writes
func dd() {
	var m = make(map[string]int)

	go func() {
		for {
			m["a"]++
		}
	}()

	go func() {
		for {
			m["a"]++
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout stop")
	}
}

// 并发使用map 加锁
func ee() {
	var m = make(map[string]int)
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)

	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()

	//go func() {
	//	for {
	//		lock.Lock()
	//		m["a"]++
	//		lock.Unlock()
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		lock.Lock()
	//		m["a"]++
	//		//fmt.Println(m["a"])
	//		lock.Unlock()
	//	}
	//}()

	select {
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}
	fmt.Println(m)

}
func ff() {
	m := make(map[string]int)
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)

	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.Lock()
			m["a"]++
			fmt.Println(m["a"])
			lock.Unlock()
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	}
}
