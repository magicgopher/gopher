package main

import "fmt"

// 通用的 Filter 函数
// 实现一个可以过滤任何类型切片的函数，只需要传入一个针对元素的类型的断言函数

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	result1 := Filter(numbers, func(n int) bool {
		// 过滤符合条件的元素
		return n%2 == 0
	})
	fmt.Println("过滤后的切片元素:", result1)

	strings := []string{"Apple", "Banana", "Cherry", "Date"}
	result2 := Filter(strings, func(s string) bool {
		// 过滤符合条件的元素
		return len(s) > 5
	})
	fmt.Println("过滤后的切片元素:", result2)
}
