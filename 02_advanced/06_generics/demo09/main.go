package main

import "fmt"

type MyInt int

type Integer interface {
	// ~ (波浪号): 这个符号表示底层类型
	// ~int 表示任何底层类型是 int 的类型。例如，如果你定义了 type MyInt int，那么 MyInt 就满足 ~int
	// 同样，~int8 表示任何底层类型是 int8 的类型，以此类推
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Add[T Integer](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("Result:", Add(2, 3))
	fmt.Println("Result:", Add(-4, -3))

	fmt.Println("==========")

	fmt.Println("Result:", Add(MyInt(2), MyInt(3)))
	fmt.Println("Result:", Add(MyInt(-1), MyInt(-2)))

	// 执行结果
	// 5
	// -7
	// 5
	// -3
}
