package main

import "fmt"

// 泛型
// go 在 1.18 版本中新加入了一个泛型的概念
// 允许在定义函数、结构体或者接口时使用类型参数（Type Parameters），而不是具体的类型。
// 这些类型参数在调用或实例化时被具体的类型替换。

func main() {
	// int类型切片
	ints := []int{1, 2, 3, 4, 5}
	r1 := IntSliceSum(ints)
	fmt.Println("Result:", r1)

	floats := []float64{1.1, 2.2, 3, 3, 4.4, 5.5}
	r2 := Float64SliceSum(floats)
	fmt.Println("Result:", r2)

	r3 := Sum(ints)
	fmt.Println("Result:", r3)
	r4 := Sum(floats)
	fmt.Println("Result:", r4)
}

// IntSliceSum int类型切片元素求和
func IntSliceSum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

// Float64SliceSum float64类型的切片元素求和
func Float64SliceSum(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	return sum
}

// ... 更多类型的 Sum 函数

// 使用泛型
// 定义一个类型约束，表示支持加法操作的类型
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// 使用类型参数 [T Number] 定义泛型函数 Sum
func Sum[T Number](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
