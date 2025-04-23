package main

import "fmt"

func main() {
	intStruct := MyGenericsStruct[int]{Data: 100}
	fmt.Printf("Type:%T, Result:%v\n", intStruct, intStruct)
	stringStruct := MyGenericsStruct[string]{Data: "hello"}
	fmt.Printf("Type:%T, Result:%v\n", stringStruct, intStruct)
}

// 结构体使用泛型
type MyGenericsStruct[T any] struct {
	Data T
}
