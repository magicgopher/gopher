package main

import (
	"fmt"
	"sync"
)

var counter int // 全局变量 counter 是临界资源

func main() {
	/*
		临界资源：是指在并发执行的多个 goroutine 中，同一时刻只能被一个 goroutine 访问和操作的共享资源。
	*/
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++ // 并发访问 counter，可能导致竞争条件
	}
}
