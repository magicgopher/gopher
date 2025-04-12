package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10       // 定义票全局变量，这个票就是临界资源
var mutex sync.Mutex  // 互斥锁
var wg sync.WaitGroup // 计数器

func main() {

	// 这4个窗口表示4个goroutine，它们会操作同一个全局变量
	wg.Add(4) // 这里4个goroutine，计数器的值就是4
	go sellTicket("售票窗口1")
	go sellTicket("售票窗口2")
	go sellTicket("售票窗口3")
	go sellTicket("售票窗口4")

	// 主 goroutine（main）睡眠4秒确保其他goroutine可以够时间运行完成
	//time.Sleep(4 * time.Second)
	wg.Wait() // 阻塞主goroutine，等待计数器为0才继续执行后续操作
}

// sellTicket 函数模拟售票
func sellTicket(window string) {
	rand.NewSource(time.Now().UnixNano())
	defer wg.Done() // 等函数执行完成，计数器-1操作
	for {
		// 上锁
		mutex.Lock()
		if ticket > 0 {
			// 模拟售票时间
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("%s正在售出:%d号票\n", window, ticket)
			// 减票
			ticket--
		} else {
			mutex.Unlock() // 释放锁
			fmt.Printf("%s票已经卖完了...\n", window)
			break // 票已经卖完
		}
		// 释放锁
		mutex.Unlock()
	}
}
