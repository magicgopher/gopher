package main

import "fmt"

// 示例：通用的 Set 结构
// 实现一个可以存储任何可比较类型元素的集合，并提供添加、删除、检查元素是否存在等操作。

type Set[T comparable] struct {
	elements map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]bool)}
}

func (s *Set[T]) Add(elem T) {
	s.elements[elem] = true
}

func (s *Set[T]) Contains(elem T) bool {
	return s.elements[elem]
}

func (s *Set[T]) Remove(elem T) {
	delete(s.elements, elem)
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func main() {
	intSet := NewSet[int]()
	intSet.Add(1)
	intSet.Add(2)
	fmt.Println("int类型的set结构中是否包含1:", intSet.Contains(1)) // 输出: set结构中是否包含1: true

	stringSet := NewSet[string]()
	stringSet.Add("apple")
	stringSet.Add("banana")
	fmt.Println("string类型的Set结构中的大小:", stringSet.Size()) // 输出: string类型的Set结构中的大小: 2
}
