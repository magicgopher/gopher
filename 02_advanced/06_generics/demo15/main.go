package main

import "fmt"

// Map - 泛型Map，支持任意类型的Key-Value
type Map[K comparable, V any] struct {
	items map[K]V
}

// NewMap - 创建一个支持任意类型Map
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		items: make(map[K]V),
	}
}

// Put - 插入或更新键值对
func (m *Map[K, V]) Put(key K, value V) {
	m.items[key] = value
}

// Get - 获取指定键的值
func (m *Map[K, V]) Get(key K) (V, bool) {
	value, exists := m.items[key]
	return value, exists
}

// Delete - 删除指定键的键值对
func (m *Map[K, V]) Delete(key K) {
	delete(m.items, key)
}

// Size - 返回 Map 中键值对的数量
func (m *Map[K, V]) Size() int {
	return len(m.items)
}

func main() {
	// 示例 1: 字符串键，整数值的 Map
	stringIntMap := NewMap[string, int]()
	stringIntMap.Put("one", 1)
	stringIntMap.Put("two", 2)
	stringIntMap.Put("three", 3)
	fmt.Println("字符串-整数 Map 大小:", stringIntMap.Size()) // 输出: 3

	// 获取值
	if value, exists := stringIntMap.Get("two"); exists {
		fmt.Println("键 'two' 的值:", value) // 输出: 2
	} else {
		fmt.Println("键 'two' 不存在")
	}

	// 删除键
	stringIntMap.Delete("one")
	fmt.Println("删除 'one' 后大小:", stringIntMap.Size()) // 输出: 2

	// 示例 2: 整数键，字符串值的 Map
	intStringMap := NewMap[int, string]()
	intStringMap.Put(1, "apple")
	intStringMap.Put(2, "banana")
	intStringMap.Put(3, "cherry")
	fmt.Println("\n整数-字符串 Map 大小:", intStringMap.Size()) // 输出: 3

	// 获取值
	if value, exists := intStringMap.Get(2); exists {
		fmt.Println("键 2 的值:", value) // 输出: banana
	} else {
		fmt.Println("键 2 不存在")
	}

	// 示例 3: 字符串键，结构体值的 Map
	type User struct {
		Name string
		Age  int
	}

	stringUserMap := NewMap[string, User]()
	stringUserMap.Put("alice", User{Name: "Alice", Age: 25})
	stringUserMap.Put("bob", User{Name: "Bob", Age: 30})
	fmt.Println("\n字符串-结构体 Map 大小:", stringUserMap.Size()) // 输出: 2

	// 获取值
	if user, exists := stringUserMap.Get("alice"); exists {
		fmt.Printf("键 'alice' 的值: %+v\n", user) // 输出: {Name:Alice Age:25}
	} else {
		fmt.Println("键 'alice' 不存在")
	}
}
