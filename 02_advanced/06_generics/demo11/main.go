package main

import "fmt"

func Print[T any](s T) {
	fmt.Println(s)
}

func main() {
	s := []int{1, 2, 3}

	// 显示指定参数类型
	Print[[]int](s)

	// 推断参数类型
	Print(s)
}
