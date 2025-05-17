package main

import "fmt"

func inSlice[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(inSlice([]string{"a", "b", "c"}, "c"))
	fmt.Println(inSlice([]string{"a", "b", "c"}, "d"))

	// 执行结果
	// 2
	// -1
}
