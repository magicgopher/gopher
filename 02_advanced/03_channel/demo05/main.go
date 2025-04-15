package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // 无缓冲 channel

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("接收数据:", <-ch) // 2秒后接收
	}()

	fmt.Println("发送数据...")
	ch <- "hello" // 发送数据，主协程在此阻塞，直到有接收方
	fmt.Println("发送完成")
}
