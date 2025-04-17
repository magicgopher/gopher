package main

import (
	"fmt"
)

func main() {
	// 创建一个无缓冲的channel
	ch1 := make(chan int)

	// 子 goroutine 在 channel 读取数据
	go func() {
		fmt.Println("正在尝试接收数据...")
		// 子 goroutine 并没有接收数据
		// 会出现：fatal error: all goroutines are asleep - deadlock!
	}()

	// 主 goroutine 向 channel 中发送数据
	fmt.Println("主 goroutine 向通道发送数据...")
	ch1 <- 66
	fmt.Println("数据发送完成...")
}
