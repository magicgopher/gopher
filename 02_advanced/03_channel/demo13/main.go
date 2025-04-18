package main

import (
	"fmt"
	"time" // 引入 time 包用于演示等待
)

// sender 函数接收一个只写的通道 (chan<- int)
// 它只能向这个通道发送数据
func sender(ch chan<- int) {
	fmt.Println("发送者: Goroutine 启动。正在发送数据...")
	ch <- 100 // 发送数据
	fmt.Println("发送者: 数据已发送。")
	// x := <-ch // 这行会编译错误: invalid operation: receive from send-only type chan<- int
	close(ch) // 发送完成后，通常由发送方关闭通道
	fmt.Println("发送者: 通道已关闭。")
}

// receiver 函数接收一个只读的通道 (<-chan int)
// 它只能从这个通道接收数据
func receiver(ch <-chan int) {
	fmt.Println("接收者: Goroutine 启动。正在等待数据...")
	// ch <- 200 // 这行会编译错误: invalid operation: send to receive-only type <-chan int

	// 使用 for range 循环可以持续接收，直到通道关闭
	// 或者使用 data, ok := <- ch 判断通道是否关闭
	data, ok := <-ch
	if ok {
		fmt.Printf("接收者: 收到数据: %d\n", data)
	} else {
		// 如果通道在发送前就关闭了，会进入这里
		fmt.Println("接收者: 通道已关闭，未收到数据。")
	}

	// 再次检查通道是否真的关闭了
	data, ok = <-ch // 如果通道已关闭，ok 会是 false
	if !ok {
		fmt.Println("接收者: 确认通道已关闭。")
	}
}

func main() {
	/*
	   单向通道 (Unidirectional Channels)
	      chan <- T : 只支持发送数据 (Send-only)
	      <- chan T : 只支持接收数据 (Receive-only)

	   主要用途：
	   1. 函数参数类型限制：明确函数对通道的操作权限，防止误操作。
	   2. 提高代码可读性和安全性。

	   转换规则：
	   - 双向通道 (chan T) 可以隐式转换为任意一种单向通道 (chan<- T 或 <-chan T)。
	   - 单向通道不能转换为双向通道。
	*/

	fmt.Println("--- 基本声明 ---")
	ch1 := make(chan int)   // 双向通道 (Bidirectional)
	ch2 := make(chan<- int) // 单向：只能发送 (Send-only)
	ch3 := make(<-chan int) // 单向：只能接收 (Receive-only)

	// 注意: %T 打印的是变量的类型，这部分保留英文是标准的 Go 类型表示
	fmt.Printf("ch1 类型: %T\n", ch1)
	fmt.Printf("ch2 类型: %T\n", ch2)
	fmt.Printf("ch3 类型: %T\n", ch3)

	// --- 实际应用演示 ---
	fmt.Println("\n--- 实际应用示例 ---")
	biDirectionalChan := make(chan int) // 创建一个双向通道

	// 启动 sender goroutine
	// 将双向通道 biDirectionalChan 传递给需要 chan<- int 的函数
	// Go 会自动进行类型转换
	go sender(biDirectionalChan)

	// 启动 receiver goroutine
	// 将同一个双向通道 biDirectionalChan 传递给需要 <-chan int 的函数
	// Go 也会自动进行类型转换
	go receiver(biDirectionalChan)

	// --- 编译时错误示例 (取消注释会报错) ---
	// sendOnlyChan := make(chan<- int)
	// <-sendOnlyChan // Error: receive from send-only type chan<- int

	// recvOnlyChan := make(<-chan int)
	// recvOnlyChan <- 1 // Error: send to receive-only type <-chan int

	// var anotherBiChan chan int = sendOnlyChan // Error: cannot use sendOnlyChan (variable of type chan<- int) as chan int value in assignment
	// var yetAnotherBiChan chan int = recvOnlyChan // Error: cannot use recvOnlyChan (variable of type <-chan int) as chan int value in assignment

	// 等待 goroutine 执行完成
	// 注意：在实际项目中，通常使用 sync.WaitGroup 来更精确地同步 goroutine，
	// 而不是依赖固定的 sleep 时间。这里为了简单起见使用 Sleep。
	fmt.Println("主程序: 等待 Goroutine 完成...")
	time.Sleep(1 * time.Second) // 等待足够的时间让 goroutine 运行

	fmt.Println("主程序: 执行完毕。")
}
