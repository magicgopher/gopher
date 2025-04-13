package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ticket = 10 // 定义票全局变量，这个票就是临界资源

func main() {

	// 创建一个用于传递票号的 channel
	// 使用无缓冲的 channel，保证了发送和接收操作的同步性
	ticketChan := make(chan int)

	// 这4个窗口表示4个goroutine，它们会操作同一个全局变量（通过 channel 间接操作）
	go sellTicket("售票窗口1", ticketChan)
	go sellTicket("售票窗口2", ticketChan)
	go sellTicket("售票窗口3", ticketChan)
	go sellTicket("售票窗口4", ticketChan)

	// 主 goroutine 负责将票号发送到 channel 中
	// 循环从初始票数递减到 1，并将每个票号发送到 ticketChan
	for i := ticket; i > 0; i-- {
		// 将当前的票号发送到 channel 中，等待售票窗口 goroutine 接收
		ticketChan <- i
	}

	// 当所有票号都发送完毕后，可以关闭 channel，通知接收方不再有新的票了
	// close(ticketChan) // 在这个例子中，不关闭 channel 也能正常结束，因为售票 goroutine 在接收不到数据时会阻塞

	// 主 goroutine（main）睡眠4秒确保其他goroutine可以够时间运行完成
	// 这里的睡眠是为了等待所有的售票 goroutine 完成它们的工作
	time.Sleep(4 * time.Second)
}

// sellTicket 函数模拟售票
func sellTicket(window string, ticketChan chan int) {
	rand.NewSource(time.Now().UnixNano())
	for {
		//
		ticket := <-ticketChan
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
