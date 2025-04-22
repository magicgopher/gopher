package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 动态指定 map 的键和值类型
	keyType := reflect.TypeOf("")  // string 类型作为键
	valueType := reflect.TypeOf(0) // int 类型作为值
	mapType := reflect.MapOf(keyType, valueType)

	// 使用 reflect.MakeMap 创建一个新的 map
	newMap := reflect.MakeMap(mapType)

	fmt.Printf("创建的 Map 类型: %s, Kind: %s\n", newMap.Type(), newMap.Kind())
	fmt.Printf("初始长度: %d\n", newMap.Len())

	// 动态向 map 中添加键值对
	keysAndValues := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	for key, value := range keysAndValues {
		// 创建键和值的 reflect.Value
		keyValue := reflect.ValueOf(key)
		valueValue := reflect.ValueOf(value)

		// 使用 SetMapIndex 设置键值对
		newMap.SetMapIndex(keyValue, valueValue)
		fmt.Printf("添加 {%q: %d} 后，长度: %d, 值: %v\n", key, value, newMap.Len(), newMap.Interface())
	}

	fmt.Printf("最终的 Map: %v (类型: %s)\n", newMap.Interface(), newMap.Type())

	// 动态获取 map 中的值
	keysToGet := []string{"apple", "grape"}
	for _, key := range keysToGet {
		keyValue := reflect.ValueOf(key)
		valueValue := newMap.MapIndex(keyValue)

		if valueValue.IsValid() {
			fmt.Printf("键 %q 的值: %v (类型: %s, Kind: %s)\n", key, valueValue.Interface(), valueValue.Type(), valueValue.Kind())
		} else {
			fmt.Printf("键 %q 不存在于 Map 中\n", key)
		}
	}

	// 动态修改 map 中的值 (需要 Value 的可寻址性)
	if newMap.CanSet() {
		// 这里通常为 false，因为 MakeMap 返回的 Value 本身不可直接 Set
	} else {
		// 要修改 map 中的值，需要获取其可寻址的 Value
		// 可以通过创建一个指向 map 的指针的 Value 来实现
		mapPtr := reflect.New(mapType)
		mapPtr.Elem().Set(newMap)
		editableMap := mapPtr.Elem()

		keyToUpdate := reflect.ValueOf("banana")
		newValue := reflect.ValueOf(20)
		editableMap.SetMapIndex(keyToUpdate, newValue)
		fmt.Printf("修改 'banana' 的值后: %v\n", editableMap.Interface())

		// 动态删除 map 中的键值对
		keyToDelete := reflect.ValueOf("cherry")
		editableMap.SetMapIndex(keyToDelete, reflect.Value{}) // 设置零值来删除
		fmt.Printf("删除 'cherry' 后: %v\n", editableMap.Interface())
	}
}
