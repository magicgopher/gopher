package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main Start")
	go sayHello() // 启动一个goroutine
	//time.Sleep(1 * time.Second) // 休眠1秒等待 goroutine 执行
	fmt.Println("Main End")
}

func sayHello() {
	fmt.Println("Hello, World!")
}
