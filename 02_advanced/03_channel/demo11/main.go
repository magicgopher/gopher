package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		关闭通道：close(ch)
			子 goroutine：发送10个数据
				每当发送一个数据，阻塞一次，主 goroutine 就接收一次，接收之后解除阻塞。

			主 goroutine：接收数据
				每当接收数据，就阻塞一次，等待 子 goroutine 发送数据，发送完数据解除阻塞。
	*/

	ch1 := make(chan int)
	go sendData(ch1)

	for {
		time.Sleep(400 * time.Millisecond)
		v, ok := <-ch1
		if !ok {
			fmt.Println("已经读取了所有数据...", ok)
			break
		}
		fmt.Println("读取的数据:", v)
	}

	fmt.Println("main...over...")
}

func sendData(ch chan int) {
	defer close(ch) // 使用 defer 关闭通道
	// 发送方：10条数据
	for i := range 10 {
		ch <- i // 将i发送到通道中
	}
}
