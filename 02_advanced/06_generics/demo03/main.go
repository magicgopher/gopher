package main

import (
	"fmt"
)

// 泛型类型

// 示例：一个泛型栈类型的定义，它可以存储任意类型的数据。

// Stack 栈
type Stack[T any] struct {
	data []T
}

// Push 将元素压入栈
func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

// Pop 从栈中弹出元素
func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		var zero T // 返回类型的零值
		return zero
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x
}

func main() {
	// 示例1：使用整数栈
	var intStack Stack[int]
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Println("整数栈弹出:", intStack.Pop()) // 输出: 3
	fmt.Println("整数栈弹出:", intStack.Pop()) // 输出: 2
	fmt.Println("整数栈弹出:", intStack.Pop()) // 输出: 1
	fmt.Println("整数栈弹出:", intStack.Pop()) // 输出: 0（零值）

	// 示例2：使用字符串栈
	var stringStack Stack[string]
	stringStack.Push("hello")
	stringStack.Push("world")
	fmt.Println("字符串栈弹出:", stringStack.Pop()) // 输出: world
	fmt.Println("字符串栈弹出:", stringStack.Pop()) // 输出: hello
	fmt.Println("字符串栈弹出:", stringStack.Pop()) // 输出: ""（零值）
}
