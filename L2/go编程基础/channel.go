package main

import (
	"fmt"
	"time"
)

/*
	gorouting 和 channel
*/

func main() {
	aa()
	//bb()
}

//仅发送数据的channel函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("发送的数据，", i)
	}
	//关闭通道
	close(ch)
}

//仅接收数据的channel函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Println("接收到的数据，", v)
	}
}
func aa() {
	//var ch chan int
	//ch := make(chan int)
	//创建一个带缓冲的channel
	ch := make(chan int, 3)

	//发送数据
	go sendOnly(ch)

	//接收数据
	go receiveOnly(ch)

	//time.Sleep(5 * time.Second)
	//使用select 多路复用

	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("channel 已关闭")
				return
			}
			fmt.Println("主gorouting接收到数据", v)
		case <-timeout:
			fmt.Println("操作超时！")
			return
		default:
			fmt.Println("没有数据，等待中！")
			time.Sleep(5 * time.Second)
		}
	}

}
