package main

import "fmt"

func main() {
	// close函数来关闭channel
	// 1. 创建一个带缓冲的 channel (容量为 2)
	// 使用缓冲 channel 可以在发送者不阻塞的情况下发送多个值，
	// 并且允许我们在同一个 goroutine 中发送、关闭、然后接收。
	ch := make(chan int, 2)

	// 2. 向 channel 发送数据
	fmt.Println("发送数据 10")
	ch <- 10
	fmt.Println("发送数据 20")
	ch <- 20

	// 3. 关闭 channel
	// 表明不会再有新的数据发送到这个 channel 了
	fmt.Println("关闭 channel...")
	close(ch)
	fmt.Println("Channel 已关闭。")

	// 4. 从 channel 接收数据
	// 即使 channel 关闭了，我们仍然可以接收里面已缓冲的数据。
	fmt.Println("开始接收数据...")

	// 接收第一个缓冲的数据
	val1, ok1 := <-ch
	fmt.Printf("接收到: %d, ok: %t\n", val1, ok1) // 输出: 接收到: 10, ok: true

	// 接收第二个缓冲的数据
	val2, ok2 := <-ch
	fmt.Printf("接收到: %d, ok: %t\n", val2, ok2) // 输出: 接收到: 20, ok: true

	// 此时 channel 已关闭，并且所有缓冲数据都已被接收完毕
	// 再次尝试接收会立即返回零值和 false
	val3, ok3 := <-ch
	fmt.Printf("再次接收 (channel 空且已关闭): %d, ok: %t\n", val3, ok3) // 输出: 再次接收...: 0, ok: false

	fmt.Println("所有数据接收完毕。")

	// 你也可以使用 for range 来接收，它会自动处理关闭的情况：
	// (如果上面的接收代码注释掉，可以取消下面代码的注释来运行)
	/*
		ch2 := make(chan string, 1)
		ch2 <- "hello"
		close(ch2) // 发送后关闭
		fmt.Println("\n使用 for range 接收:")
		for msg := range ch2 {
			fmt.Printf("For range 接收到: %s\n", msg) // 输出: For range 接收到: hello
		}
		fmt.Println("For range 结束 (因为 channel 关闭且数据接收完)")
	*/
}
