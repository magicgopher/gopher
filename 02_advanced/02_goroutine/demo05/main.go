package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 票
var ticket = 100

// 计数器
var wg sync.WaitGroup

// 互斥锁
var mutex sync.Mutex

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	numWindows := 5
	wg.Add(numWindows)
	fmt.Printf("初始票数: %d\n", ticket)

	for i := 1; i <= numWindows; i++ {
		go sellTicket(i)
	}

	wg.Wait()
	fmt.Println("售票结束")
	fmt.Printf("最终剩余票数: %d (预期为 0，但很可能不正确)\n", ticket)
}

// sellTicket 卖票操作 没有使用锁会有临界资源安全问题
//func sellTicket(window int) {
//	for ticket > 0 {
//		time.Sleep(20 * time.Millisecond)
//		ticket--
//		fmt.Printf("%d窗口正在售卖编号%d的票,剩余票数:%d\n", window, ticket+1, ticket)
//	}
//	wg.Done()
//}

// sellTicket 卖票操作 使用锁
func sellTicket(window int) {
	mutex.Lock() // 上锁
	for ticket > 0 {
		time.Sleep(20 * time.Millisecond)
		ticket--
		fmt.Printf("%d窗口正在售卖编号%d的票,剩余票数:%d\n", window, ticket+1, ticket)
	}
	mutex.Unlock() // 释放锁
	wg.Done()
}
