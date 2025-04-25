package main

import "fmt"

// 通用的数据处理管道
// 在数据处理流程中，经常需要对数据进行一系列的操作。
// 例如映射 (map)、过滤 (filter)、归约 (reduce) 等。泛型可以帮助我们构建通用的处理步骤。

func Map[T any, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = transform(item)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3}
	result1 := Map(numbers, func(n int) int {
		return n * n
	})
	fmt.Println(result1)
}
