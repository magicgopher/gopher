package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓冲的channel
	ch1 := make(chan string)

	go func() {
		fmt.Println("子 goroutine 开始等待接收数据...")
		data := <-ch1
		fmt.Println("子 goroutine 接收到的数据:", data)
	}()

	// 主 goroutine 等待一段时间，但不发送数据
	fmt.Println("主 goroutine 等待3秒")
	time.Sleep(3 * time.Second)
	fmt.Println("主协程：3 秒已过，仍然不发送数据")

	// 为了避免程序直接退出，添加一个长时间等待
	fmt.Println("主协程：继续运行，子协程仍在阻塞")
	time.Sleep(2 * time.Second)
	fmt.Println("主协程：程序结束")
}
