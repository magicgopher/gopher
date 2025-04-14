package main

import "fmt"

func main() {
	/*
		channel,通道

		go建议使用channel通道来实现goroutine之间的通信

		go社区有句经典的话：不要通过共享内存来通信，应该通过通信来共享内存。
	*/

	var a chan int
	fmt.Printf("%T, %v\n", a, a)

	if a == nil {
		fmt.Println("channel通道是nil的，不能直接使用，需要先创建通道")
		a = make(chan int) // 通过make函数创建通道
		fmt.Println(a)
	}
	test1(a)
}

func test1(ch chan int) {
	fmt.Printf("%T, %v\n", ch, ch) // 打印传入进来的通道类型参数，这里输出的是地址值

}
