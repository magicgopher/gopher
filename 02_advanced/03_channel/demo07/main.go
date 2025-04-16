package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // 无缓冲 channel

	go func() {
		time.Sleep(2 * time.Second) // 模拟延迟
		fmt.Println("发送数据到 channel")
		ch <- "Hello" // 发送数据，阻塞直到有接收者
	}()

	fmt.Println("准备接收数据")
	data := <-ch // 接收数据，阻塞直到有发送者
	fmt.Println("收到数据:", data)
}
