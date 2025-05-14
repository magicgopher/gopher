package main

import "fmt"

// GenericMap 定义泛型 Map 类型
type GenericMap[K comparable, V int | string | byte] map[K]V

// Add 向 map 中添加键值对
func (m GenericMap[K, V]) Add(key K, value V) {
	m[key] = value
}

// Get 根据key获取对应的value
func (m GenericMap[K, V]) Get(key K) (V, bool) {
	value, ok := m[key]
	return value, ok
}

// Delete 方法：删除指定键
func (m GenericMap[K, V]) Delete(key K) {
	delete(m, key)
}

// Len 方法：返回 map 的长度
func (m GenericMap[K, V]) Len() int {
	return len(m)
}

func main() {
	// 示例 1：使用 string 作为键，int 作为值
	intMap := make(GenericMap[string, int])
	intMap.Add("one", 1)
	intMap.Add("two", 2)
	intMap.Add("three", 3)

	fmt.Println("String-Int Map:")
	fmt.Printf("Length: %d\n", intMap.Len())
	if value, ok := intMap.Get("two"); ok {
		fmt.Printf("Key 'two' value: %d\n", value)
	} else {
		fmt.Println("Key 'two' not found")
	}

	intMap.Delete("one")
	fmt.Printf("After deleting 'one', Length: %d\n", intMap.Len())
	fmt.Println("Map contents:", intMap)
	fmt.Println()

	// 示例 2：使用 int 作为键，string 作为值
	stringMap := make(GenericMap[int, string])
	stringMap.Add(1, "apple")
	stringMap.Add(2, "banana")
	stringMap.Add(3, "cherry")

	fmt.Println("Int-String Map:")
	fmt.Printf("Length: %d\n", stringMap.Len())
	if value, ok := stringMap.Get(2); ok {
		fmt.Printf("Key 2 value: %s\n", value)
	} else {
		fmt.Println("Key 2 not found")
	}
	fmt.Println("Map contents:", stringMap)
	fmt.Println()

	// 示例 3：使用 string 作为键，byte 作为值
	byteMap := make(GenericMap[string, byte])
	byteMap.Add("a", 'A')
	byteMap.Add("b", 'B')

	fmt.Println("String-Byte Map:")
	fmt.Printf("Length: %d\n", byteMap.Len())
	if value, ok := byteMap.Get("a"); ok {
		fmt.Printf("Key 'a' value: %c\n", value)
	} else {
		fmt.Println("Key 'a' not found")
	}
	fmt.Println("Map contents:", byteMap)
}
