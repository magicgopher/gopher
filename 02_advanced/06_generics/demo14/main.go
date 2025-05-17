package main

import "fmt"

// 泛型实现数据结构，简单示例，实现一个队列

// Queue 队列
type Queue[T any] struct {
	items []T
}

// Put 将数据放入队列尾部
func (q *Queue[T]) Put(value T) {
	q.items = append(q.items, value)
}

// Pop 从队列头部取出并从头部删除对应数据
func (q *Queue[T]) Pop() (T, bool) {
	var value T
	// 队列没有数据
	if len(q.items) == 0 {
		return value, true
	}
	// 队列有数据
	value = q.items[0]    // 取出队列第一个元素
	q.items = q.items[1:] // 将剩余的元素赋值给items
	return value, len(q.items) == 0
}

// Size 队列大小
func (q Queue[T]) Size() int {
	return len(q.items)
}

func main() {
	// 示例1：整数队列
	intQueue := Queue[int]{}
	fmt.Printf("整数队列初始化大小:%d\n", intQueue.Size())

	// 向队列添加元素
	intQueue.Put(100)
	intQueue.Put(200)
	intQueue.Put(300)
	fmt.Printf("添加3个元素后大小:%d\n", intQueue.Size())

	// 取出队列的元素
	for {
		value, isEmpty := intQueue.Pop()
		if isEmpty {
			fmt.Println("整数队列已空。")
			break
		}
		fmt.Println("整数队列取出的元素:", value)
	}

	fmt.Println("==========")

	// 示例 2: 使用字符串队列
	stringQueue := Queue[string]{}
	stringQueue.Put("apple")
	stringQueue.Put("banana")
	stringQueue.Put("cherry")
	fmt.Println("字符串队列大小:", stringQueue.Size()) // 输出: 3

	// 取出并打印元素
	for {
		value, isEmpty := stringQueue.Pop()
		if isEmpty {
			fmt.Println("字符串队列已空")
			break
		}
		fmt.Println("取出元素:", value)
	}
}
