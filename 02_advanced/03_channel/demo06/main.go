package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个 channel 缓冲区大小为5
	ch1 := make(chan int, 5)
	fmt.Printf("类型:%T, 长度:%d, 容量:%d\n", ch1, len(ch1), cap(ch1))

	// 子 goroutine
	go func() {
		time.Sleep(1 * time.Second)
		for i := 1; i <= 10; i++ {
			data := <-ch1
			fmt.Println("子 goroutine 接收数据, data:", data)
		}
		fmt.Println("子 goroutine 结束...")
	}()

	// 主 goroutine 发送数据
	for i := 1; i <= 10; i++ {
		fmt.Println("主 goroutine 发送数据:", i)
		ch1 <- i
	}
	time.Sleep(3 * time.Second)
	fmt.Println("主 goroutine 结束...")
}
