package main

// Add 一个简单的求和函数
//func Add(a,b int) int {
//	return a + b
//}

// AddFloat 新增一个函数用来支持其他类型的求和函数
//func AddFloat(a,b float64) float64 {
//	return a + b
//}

// Add 使用空接口+反射的方式来支持其他类型的求和函数
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

// Add 使用新特性泛型来支持其他类型的求和函数
func Add[T int | float32 | float64](a, b T) T {
	return a + b
}

func main() {
	//fmt.Println(Add(2, 3))
	//fmt.Println(Add(1.1, 2.3))
}
