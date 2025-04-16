package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个有缓冲的通道
	ch1 := make(chan int, 5)

	// 子 goroutine 接收数据
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 10; i++ {
			data := <-ch1 // 接收数据
			fmt.Println("子 goroutine 接收到的数据:", data)
		}
		fmt.Println("子 goroutine 结束...")
	}()

	// 主 goroutine 发送数据
	for i := 0; i < 10; i++ {
		fmt.Println("主 goroutine 发送数据, i:", i)
		ch1 <- i
	}
	fmt.Println("主 goroutine 结束...")

	// 主 goroutine 睡眠2秒，等待子 goroutine 执行完成
	time.Sleep(2 * time.Second)
}
