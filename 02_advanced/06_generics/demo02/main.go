package main

import "fmt"

// 泛型函数

// 示例：一个简单的泛型函数，它可以接受任意类型的参数，并返回一个切片

// toSlice 返回一个切片
func toSlice[T any](args ...T) []T { // T 表示任意类型，args 表示可变参数
	return args
}

func main() {
	fmt.Println(toSlice("go", "java", "python")) // []string{"go", "java", "python"}
	fmt.Println(toSlice(1, 2, 3))                // []int{1,2,3}
}
