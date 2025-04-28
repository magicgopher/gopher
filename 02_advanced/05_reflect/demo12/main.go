package main

import (
	"fmt"
	"reflect"
)

// 假设我们有一个函数，它接收一个任意类型的输入，并且我们想要打印出这个输入的类型和值。

// 使用普通变量的方式
// 如果我们事先知道输入的类型，我们可以为每种可能的类型编写一个函数。
// 这种方式的缺点是，如果我们需要处理更多的类型，就需要编写更多的 print 函数，代码会变得冗余且难以维护。
//func main() {
//	var num int = 10
//	var text string = "hello"
//	printInt(num)
//	printString(text)
//}
//
//func printInt(val int) {
//	fmt.Printf("类型: int, 值: %d\n", val)
//}
//
//func printString(val string) {
//	fmt.Printf("类型: string, 值: %s\n", val)
//}

// 使用 interface{} 和 类型断言的方式
// 我们可以使用空接口 interface{} 来接收任意类型，然后使用类型断言来判断实际类型。
// 这种方式比第一种好一些，可以处理多种类型，但仍然需要在 switch 语句中显式地列出所有我们想要处理的类型。
// 如果类型很多或者我们无法预知所有可能的类型，这种方式也会显得不够灵活。
//func main() {
//	var num int = 10
//	var text string = "hello"
//	var flag bool = true
//	printAny(num)
//	printAny(text)
//	printAny(flag)
//}
//
//func printAny(val interface{}) {
//	switch v := val.(type) {
//	case int:
//		fmt.Printf("类型: int, 值: %d\n", v)
//	case string:
//		fmt.Printf("类型: string, 值: %s\n", v)
//	default:
//		fmt.Printf("类型: 不存在, 值: %v\n", v)
//	}
//}

// 使用反射的方式
func main() {
	var num int = 10
	var text string = "hello"
	var flag bool = true
	myStruct := struct{ Name string }{Name: "Bob"}

	printWithTypeAndValue(num)
	printWithTypeAndValue(text)
	printWithTypeAndValue(flag)
	printWithTypeAndValue(myStruct)
}

func printWithTypeAndValue(val interface{}) {
	// 获取变量的类型信息
	t := reflect.TypeOf(val)
	fmt.Printf("类型: %v\n", t)

	// 获取变量的值信息
	v := reflect.ValueOf(val)
	fmt.Printf("值: %v\n", v)
}
