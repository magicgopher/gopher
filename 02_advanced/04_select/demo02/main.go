package main

import (
	"fmt"
	"time"
)

func main() {
	// select示例
	// 定义两个无缓冲通道 ch1 和 ch2 用于接收字符串消息

	// 创建两个无缓冲channel
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个 goroutine 分别向 ch1 和 ch2 发送消息
	go func() {
		time.Sleep(time.Second)
		ch1 <- "ch1:消息内容"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "ch2:消息内容"
	}()

	// 使用 select 监听两个通道
	for i := 0; i < 2; i++ {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				fmt.Println("ch1通道关闭.")
			}
			fmt.Println("ch1通道的msg:", msg1)
		case msg2, ok := <-ch2:
			if !ok {
				fmt.Println("ch2通道关闭.")
			}
			fmt.Println("ch2通道的msg:", msg2)
		}
	}

	fmt.Println("main...over...")
}
