package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 动态指定切片的元素类型为 int
	sliceType := reflect.SliceOf(reflect.TypeOf(0))

	// 使用 reflect.MakeSlice 创建一个新的切片
	newSlice := reflect.MakeSlice(sliceType, 0, 5) // 长度为 0，容量为 5

	fmt.Printf("创建的切片类型: %s, Kind: %s\n", newSlice.Type(), newSlice.Kind())
	fmt.Printf("初始长度: %d, 容量: %d\n", newSlice.Len(), newSlice.Cap())

	// 动态添加元素到切片
	elementsToAdd := []int{10, 20, 30}
	for _, val := range elementsToAdd {
		// 创建要添加的元素的 reflect.Value
		elemValue := reflect.ValueOf(val)

		// 使用 Append 函数添加元素
		newSlice = reflect.Append(newSlice, elemValue)

		fmt.Printf("添加元素 %v 后，长度: %d, 容量: %d, 值: %v\n", val, newSlice.Len(), newSlice.Cap(), newSlice.Interface())
	}

	fmt.Printf("最终的切片: %v (类型: %s)\n", newSlice.Interface(), newSlice.Type())

	// 动态获取切片中的元素
	for i := 0; i < newSlice.Len(); i++ {
		elem := newSlice.Index(i)
		fmt.Printf("索引 %d 的元素: %v (类型: %s, Kind: %s)\n", i, elem.Interface(), elem.Type(), elem.Kind())
	}

	// 动态修改切片中的元素 (需要 Value 的可寻址性)
	if newSlice.CanSet() {
		// 这里会 panic，因为 MakeSlice 返回的 Value 通常不可直接 Set
	} else {
		// 要修改切片元素，需要获取其可寻址的 Value
		// 可以通过创建一个新的可寻址的 Value 并复制元素来实现
		editableSlice := reflect.New(sliceType).Elem()
		editableSlice = reflect.AppendSlice(editableSlice, newSlice)

		if editableSlice.CanSet() { // 仍然是 false
			// editableSlice 的元素本身是不可寻址的
		} else {
			// 需要进一步获取元素的指针
			if editableSlice.Len() > 0 {
				firstElemPtr := editableSlice.Index(0).Addr()
				if firstElemPtr.CanSet() { // 仍然是 false
					// int 类型的 Value 是不可寻址的
				} else {
					// 要修改基本类型的值，需要通过指针指向该值
					firstElemValue := editableSlice.Index(0)
					if firstElemValue.Kind() == reflect.Int {
						newValue := reflect.ValueOf(100)
						// 注意：这里不能直接 Set，因为 firstElemValue 不可寻址
						// 需要通过创建一个可寻址的临时 Value 并 Set
						temp := reflect.New(firstElemValue.Type()).Elem()
						temp.Set(newValue)
						editableSlice.Index(0).Set(temp)
						fmt.Printf("尝试修改后的可编辑切片: %v\n", editableSlice.Interface())
					}
				}
			}
		}
	}
}
