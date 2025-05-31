package main

import (
	"fmt"
	"reflect"
)

type Slice[T interface{}] []T

func main() {
	var s1 Slice[int]
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("切片存储的元素数据类型是:%v\n", reflect.TypeOf(s1).Elem())
	fmt.Println(s1)

	fmt.Println("===============")

	var s2 Slice[string]
	s2 = append(s2, "go", "java", "python")
	fmt.Printf("切片存储的元素数据类型是:%v\n", reflect.TypeOf(s2).Elem())
	fmt.Println(s2)

	fmt.Println("===============")

	var s3 Slice[interface{}]
	s3 = append(s3, 1, 2.2, '3', "hello")
	fmt.Println(s3)
}
