package main

import "fmt"

func main() {
	/*
		双向通道
			chan T
				chan <- data 发送数据，写出
				data <- chan 接收数据，读取
	*/
	ch1 := make(chan string)
	done := make(chan bool)
	go sendData(ch1, done)

	// 主 goroutine 接收数据
	data := <-ch1
	fmt.Println("子goroutine发送来的数据:", data)

	// 主 goroutine 发送数据
	ch1 <- "Hi"
	<-done // 接收信号通知
	fmt.Println("main...over...")
}

func sendData(ch chan string, done chan bool) {
	ch <- "Hello" // 向通道发送数据

	data := <-ch // 从通道接收数据
	fmt.Println("主 goroutine 发送来的数据:", data)

	done <- true // 信号通知
}
