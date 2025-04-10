package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		示例：羊统计五次、牛统计三次

		goroutine的运行直接和java的Thread不所不同。

		在 Go 语言中，情况与 Java 有所不同。如果 main 函数执行完成。
		而 goroutine 还没有执行完成，那么 goroutine 不会继续执行，整个程序会立即退出。
	*/
	var wg sync.WaitGroup
	wg.Add(2) // 现在是需要统计羊五次、牛统计三次，一共是两个任务，所以计数器为2
	go func() {
		defer wg.Done() // goroutine 结束时计数器减 1
		count(5, "🐑")
	}()
	go func() {
		defer wg.Done() // goroutine 结束时计数器减 1
		count(3, "🐂")
	}()
	wg.Wait() // 等待计数器归零
	//time.Sleep(time.Second * 5) // 通过睡眠1秒来延迟main函数的退出
}

func count(n int, animal string) {
	fmt.Println("goroutine 开始运行")
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
	}
	time.Sleep(time.Millisecond * 500) // 模拟任务耗时
	fmt.Println("goroutine 执行完成")
}
