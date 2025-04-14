package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建通道
	ch1 := make(chan int)
	// 子goroutine发送数据
	go func() {
		fmt.Println("子goroutine开始执行...")
		ch1 <- 66 // 子goroutine向通道发送数据
	}()
	time.Sleep(2 * time.Second) // 主goroutine睡眠2秒
	// 主goroutine从通道接收数据
	data, ok := <-ch1
	if ok {
		fmt.Println("成功从通道接收到数据...")
		fmt.Println("data:", data)
	} else {
		fmt.Println("通道已经关闭")
	}
	fmt.Println("main...over...")
}
