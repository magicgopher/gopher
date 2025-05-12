package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 函数的反射
	f1 := fun1
	f1Value := reflect.ValueOf(f1)
	fmt.Printf("Kind: %s, Type: %s\n", f1Value.Kind(), f1Value.Type())

	f2 := fun2
	f2Value := reflect.ValueOf(f2)
	fmt.Printf("Kind: %s, Type: %s\n", f2Value.Kind(), f2Value.Type())

	// 通过反射调用函数
	f1Value.Call(nil) // 没有参数
	f2Value.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf("hello")})
}

func fun1() {
	fmt.Println("fun1()，没有参数.")
}

func fun2(i int, s string) {
	fmt.Println("fun2()，有参数.")
}
