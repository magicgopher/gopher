package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ticket = 10 // 定义票全局变量，这个票就是临界资源

func main() {
	// 这4个窗口表示4个goroutine，它们会操作同一个全局变量
	go sellTicket("售票窗口1")
	go sellTicket("售票窗口2")
	go sellTicket("售票窗口3")
	go sellTicket("售票窗口4")

	// 主 goroutine（main）睡眠4秒确保其他goroutine可以够时间运行完成
	time.Sleep(4 * time.Second)
}

// sellTicket 函数模拟售票
func sellTicket(window string) {
	rand.NewSource(time.Now().UnixNano())
	for {
		if ticket > 0 {
			// 模拟售票时间
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("%s正在售出:%d号票\n", window, ticket)
			// 减票
			ticket--
		} else {
			fmt.Printf("%s票已经卖完了...\n", window)
			break // 票已经卖完
		}
	}
}
