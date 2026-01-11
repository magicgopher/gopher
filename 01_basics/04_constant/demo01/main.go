package main

import "fmt"

func main() {
	// 常量：计算机中的一小块内存，用于存储数据，在程序运行的过程中是不可变的。

	// 方式一
	// 显式定义
	const PATH string = "www.baidu.com"
	fmt.Println(PATH)

	// 方式二
	// 隐式定义常量
	const PI = 3.14
	fmt.Println(PI)

	// 尝试修改常量的内容
	//PATH = "www.google.com" // cannot assign to PATH
	//fmt.Println(PATH)

	// 定义一组常量
	const C1, C2, C3 = 100, 3.14, "你好" // 写在一行的
	fmt.Println(C1, C2, C3)
	const ( // 使用 const () 定义一组常量
		DAY   = 1
		WEEK  = 2
		MONTH = 3
		YEAR  = 4
	)
	fmt.Println(DAY, WEEK, MONTH, YEAR)

	// 一组常量中，如果有某个常量没有赋值，默认和上一行一样的数据
	const (
		A = 100
		B
		C = "Hello"
		D
	)
	fmt.Println(A, B, C, D)
}
