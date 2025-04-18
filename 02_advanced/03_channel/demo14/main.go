package main

import (
	"fmt"
)

func main() {
	/*
		单向通道
			chan <- T 只支持发送数据（写）
			<- chan T 只支持接收数据（读）

		在实际开发中单向通道通常都是应用就是在函数或方法的入参上。
	*/

	ch1 := make(chan int) // 双向通道，支持读写
	//ch2 := make(chan<- int) // 单向通道，只支持写，不支持读
	//ch3 := make(<-chan int) // 单向通道，只支持读，不支持写

	//ch1 <- 100    // 发送数据
	//data := <-ch1 // 接收数据

	//ch2 <- 200
	//data := <-ch2 // invalid operation: cannot receive from send-only channel ch2 (variable of type chan<- int)

	//data := <-ch3
	//ch3 <- 300 // invalid operation: cannot send to receive-only channel ch3 (variable of type <-chan int)

	// 调用fun1函数可以传递ch1和ch2
	go fun1(ch1) // 这里可以传入ch1，但是受到函数的限制，这里只能是发送数据的通道
	//fun1(ch2)    // 这里可以传入ch2

	data := <-ch1
	fmt.Println("接收fun1函数发送的数据:", data)

	// 调用fun2函数可以传递ch1和ch3
	go fun2(ch1)
	ch1 <- 200
	fmt.Println("main...over...")
}

// fun1 该函数，只能操作发送数据的通道
func fun1(ch chan<- int) {
	// 在函数内部，ch1虽然是双向通道，但是这里搜到函数参数类型的限制，这里只能是单向通道，只能发送数据，不能接收数据
	ch <- 100 // 向通道发送数据
	fmt.Println("fun1函数结束...")
}

// fun2 该函数，只能接收数据的通道
func fun2(ch <-chan int) {
	data := <-ch
	fmt.Println("fun2函数，从ch通道中读取到的数据是:", data)
	fmt.Println("fun2函数结束...")
}
