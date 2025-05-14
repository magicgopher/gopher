package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 1.23456
	fmt.Println("num的值:", num)

	// 需要操作指针
	// 通过 reflect.ValueOf 获取 num 中的 reflect.Value，注意，参数必须是指针才能修改其值
	numPointer := reflect.ValueOf(&num)
	elem := numPointer.Elem()
	fmt.Println("类型:", elem.Type())       // float64
	fmt.Println("是否可以修改:", elem.CanSet()) // true

	// 重新赋值
	elem.SetFloat(9.87654)
	fmt.Println("num的新值:", num)

	// 如果reflect.ValueOf的参数不是指针，会如何？

	// 参试直接修改
	//value := reflect.ValueOf(num)
	//value.SetFloat(4.5678) // panic: reflect: reflect.Value.SetFloat using unaddressable value
	//fmt.Println(value.CanSet()) // false

	//pointer := reflect.ValueOf(num)
	//elem := pointer.Elem() // 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”
}
