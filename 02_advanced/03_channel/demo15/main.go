package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		通道的遍历
	*/
	ch1 := make(chan int)
	// 启动子 goroutine 向通道发送数据
	go sendData(ch1)

	// for...range 从通道接收值，直到通道关闭
	for v := range ch1 {
		fmt.Println("从通道中接收数据:", v)
	}
	fmt.Println("main...over...")
}

// sendData 发送数据
func sendData(ch chan int) {
	defer close(ch) // 函数结束后defer关闭通道
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- i // 向通道发送数据i
	}
}
