package main

import (
	"fmt"
	"time"
)

// select是go中的一个选择语句，采用select/case结构，类似流程控制的switch/case语句
// select只能用于监听channel的读写（接收/发送）操作
// channel用于goroutine的通信，那么select则用于监听多个channel的读写（接收/发送）操作

func main() {
	//ch1 := make(chan string)
	//ch2 := make(chan string)

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 2)

	ch1 <- "Hello select1"
	ch2 <- "Hello select2"

	go func() {
		fmt.Println("select start")
		for {
			// 多个case同时满足时，随机选择一个case执行
			select {
			// 监听ch1读操作，如果ch1有数据ok为true
			case msg1, ok := <-ch1:
				if !ok {
					fmt.Println("ch1 closed")
					return
				}
				fmt.Println(msg1)
			case msg2, ok := <-ch2:
				if !ok {
					fmt.Println("ch2 closed")
					return
				}
				fmt.Println(msg2)
				//default: // 如果case没有命中（一直阻塞）就会进入default
				//	fmt.Println("select default")
			}
		}
		fmt.Println("select end")
	}()

	//close(ch1)
	ch1 <- "Hello select1"
	ch2 <- "Hello select2"

	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
