package main

import "fmt"

// 匿名函数不支持泛型
// 日常开发中会经常的使用匿名函数

func main() {

	// 定义了一个匿名函数，将其赋值给了fn
	fn := func(a, b int) int {
		return a + b
	}

	// 错误，匿名函数不能自己定义类型实参
	//fn := func[T int | float32](a, b T) T {
	//	return a + b
	//}

	fmt.Println(fn(1, 2)) // 输出：3

}

func MyFunc[T int | float32 | float64](a, b T) {
	// 匿名函数可使用已经定义好的类型形参
	fn2 := func(i T, j T) T {
		return i + j
	}
	fn2(a, b)
}
