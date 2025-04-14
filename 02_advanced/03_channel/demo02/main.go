package main

import "fmt"

func main() {
	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中, i:", i)
		}
		// 循环结束后，向通道中写数据，表示结束了。
		ch1 <- true
		fmt.Println("结束...")
	}()

	// 这里的data是在通道中获取数据的，要是主goroutine先抢到执行的资源
	// 由于子goroutine没有向通道中写入数据，所以主goroutine是阻塞的。
	data := <-ch1 // 从通道中读取数据
	fmt.Println("main...data->", data)
	fmt.Println("main...over...")
}
