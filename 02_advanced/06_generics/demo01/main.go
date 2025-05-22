package main

import "fmt"

// Add 求和
//func Add(a, b int) int {
//	return a + b
//}

// AddFloat 求和支持float64类型
//func AddFloat(a, b float64) float64 {
//	return a + b
//}

// Add 求和函数 使用反射来实现支持任意类型参数
//func Add(a, b interface{}) interface{} {
//	switch a.(type) {
//	case int:
//		return a.(int) + b.(int)
//	case float64:
//		return a.(float64) + b.(float64)
//	default:
//		return nil
//	}
//}

// Add 求和函数 使用泛型来支持任意类型参数
func Add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(10, 20)) // 结果：30
}
