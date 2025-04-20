package main

import (
	"fmt"
	"reflect"
)

func main() {
	// reflect 反射
	// 定义：在计算机科学中，反射是指计算机程序在运行时（Run time）可以访问、检测和修改它本身状态或行为的一种能力。
	// 用比喻来说，反射就是程序在运行的时候能够“观察”并且修改自己的行为。

	// 定义一个float64类型变量
	var x = 42

	t := reflect.TypeOf(x)  // 获取类型
	v := reflect.ValueOf(x) // 获取值

	fmt.Println("Type:", t)        // Type: int
	fmt.Println("Value:", v)       // Value: 42
	fmt.Println("Kind:", t.Kind()) // Kind: int
}
