package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 55
	fmt.Printf("类型:%T, 值:%d\n", x, x)
	v := reflect.ValueOf(&x).Elem()   // 获取指针指向的值
	v.SetInt(100)                     // 修改值
	fmt.Printf("类型:%T, 值:%d\n", x, x) // 新值 100
}
