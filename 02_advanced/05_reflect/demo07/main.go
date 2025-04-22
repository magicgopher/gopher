package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明一个 int 变量并获取其指针
	number := 42
	ptr := &number

	// 使用 reflect.ValueOf 获取指针的 reflect.Value
	ptrValue := reflect.ValueOf(ptr)

	fmt.Println("--- 处理指针 ---")
	fmt.Printf("Value 类型: %s, Kind: %s\n", ptrValue.Type(), ptrValue.Kind()) // *int, ptr
	fmt.Printf("CanSet: %t\n", ptrValue.CanSet())                            // false (指针本身不可直接修改)

	// 使用 Elem() 获取指针指向的值的 reflect.Value
	elemValue := ptrValue.Elem()

	fmt.Printf("\n--- 处理指针指向的值 ---")
	fmt.Printf("Elem() 类型: %s, Kind: %s\n", elemValue.Type(), elemValue.Kind()) // int, int
	fmt.Printf("CanSet (Elem()): %t\n", elemValue.CanSet())                     // true (如果原始变量可寻址)
	fmt.Printf("原始指针指向的值: %v\n", elemValue.Interface())

	// 修改指针指向的值 (通过 Elem() 的 Set 方法)
	if elemValue.CanSet() {
		newValue := reflect.ValueOf(100)
		if elemValue.Type() == newValue.Type() {
			elemValue.Set(newValue)
			fmt.Printf("修改后指针指向的值: %v (原始变量 number: %v)\n", elemValue.Interface(), number)
		} else {
			fmt.Println("尝试设置的值类型不匹配")
		}
	} else {
		fmt.Println("指针指向的值不可修改 (原始变量不可寻址)")
	}

	// 创建一个新的指针类型的 reflect.Value
	intPtrType := reflect.PtrTo(reflect.TypeOf(0)) // *int 类型
	newPtrValue := reflect.New(intPtrType.Elem())  // reflect.New 返回指向新分配的零值的指针

	fmt.Println("\n--- 创建新的指针 ---")
	fmt.Printf("新的指针 Value 类型: %s, Kind: %s\n", newPtrValue.Type(), newPtrValue.Kind()) // *int, ptr
	fmt.Printf("新的指针指向的值 (初始): %v\n", newPtrValue.Elem().Interface())

	// 设置新指针指向的值
	newElemValue := newPtrValue.Elem()
	if newElemValue.CanSet() {
		newElemValue.Set(reflect.ValueOf(200))
		fmt.Printf("新的指针指向的值 (修改后): %v\n", newElemValue.Interface())
		fmt.Printf("新的指针指向的地址: %v\n", newPtrValue.Interface())
		fmt.Printf("新的指针指向的实际值 (通过解引用): %v\n", *newPtrValue.Interface().(*int))
	}

	// 处理 nil 指针
	var nilPtr *int
	nilPtrValue := reflect.ValueOf(nilPtr)

	fmt.Println("\n--- 处理 nil 指针 ---")
	fmt.Printf("nil 指针 Value 类型: %s, Kind: %s\n", nilPtrValue.Type(), nilPtrValue.Kind()) // *int, ptr
	fmt.Printf("IsNil(): %t\n", nilPtrValue.IsNil())                                      // true

	// 尝试对 nil 指针调用 Elem() 会 panic
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("尝试对 nil 指针调用 Elem() 导致 panic:", r)
	// 	}
	// }()
	// _ = nilPtrValue.Elem() // 这行会 panic

	// 使用 IsNil() 检查指针是否为 nil
	if !nilPtrValue.IsNil() {
		_ = nilPtrValue.Elem()
	}
}
