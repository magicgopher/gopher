package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		示例：启动多个goroutine
		一个goroutine打印数字1-5
		另一个goroutine打印字母a-e
	*/
	fmt.Println("main start")
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond) // 睡眠3秒让main函数延迟结束
	fmt.Println("main over")
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("数字:%c\n", i)
	}
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("字母:%d\n", i)
	}
}
