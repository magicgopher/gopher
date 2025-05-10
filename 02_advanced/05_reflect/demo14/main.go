package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 1.2345678910

	pointer := reflect.ValueOf(&num) // 指向 num 的指针
	value := reflect.ValueOf(num)    // num 值的副本

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接 panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是 *float64，一个是 float64，如果弄混，则会 panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer) // 指针存储的地址值
	fmt.Println(convertValue)   // 变量存储的数值副本
}
