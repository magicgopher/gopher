package main

import "fmt"

// 接口约束

// 示例：一个泛型类型，它要求其泛型类型参数实现了 fmt.Stringer 接口

// CustomType 是一个实现 fmt.Stringer 接口的自定义类型
type CustomType struct {
	value string
}

// String 实现 fmt.Stringer 接口
func (c CustomType) String() string {
	return c.value
}

// ErrorWrapper 包装 error 类型，使其实现 fmt.Stringer
type ErrorWrapper struct {
	err error
}

// String 实现 fmt.Stringer 接口
func (e ErrorWrapper) String() string {
	return e.err.Error()
}

// MyType 是一个泛型结构体，约束 T 必须实现 fmt.Stringer 接口
type MyType[T fmt.Stringer] struct {
	data T
}

// String 调用底层类型的 String 方法
func (m *MyType[T]) String() string {
	return m.data.String()
}

func main() {
	// 示例1：使用 CustomType 作为泛型参数
	custom := CustomType{value: "Hello, Generics!"}
	myType1 := MyType[CustomType]{data: custom}
	fmt.Println(myType1.String()) // 输出: Hello, Generics!

	// 示例2：使用 ErrorWrapper 包装 error 类型
	err := fmt.Errorf("this is an error")
	wrappedErr := ErrorWrapper{err: err}
	myType2 := MyType[ErrorWrapper]{data: wrappedErr}
	fmt.Println(myType2.String()) // 输出: this is an error
}
