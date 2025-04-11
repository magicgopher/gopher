package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 模拟临界资源
var counter int

// 计数器
var wg sync.WaitGroup

// 互斥锁
//var lock sync.Mutex

func main() {
	/*
		临界资源：是指在并发执行的多个 goroutine 中，同一时刻只能被一个 goroutine 访问和操作的共享资源。
	*/

	runtime.GOMAXPROCS(runtime.NumCPU()) // 利用所有可用的 CPU 核心
	numGoroutines := 10

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go incrementCounter()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", counter)
}

func incrementCounter() {
	for i := 0; i < 10000; i++ {
		//lock.Lock()   // 获取锁，阻止其他 goroutine 访问临界区
		counter++ // 这是一个临界资源，多个 goroutine 同时访问和修改
		//lock.Unlock() // 释放锁，允许其他 goroutine 访问
	}
	wg.Done()
}
