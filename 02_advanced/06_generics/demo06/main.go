package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	// int类型参数
	res := Sum(10, 20)
	fmt.Println(res)

	// float64类型参数
	res2 := SumFloat64(11.22, 33.44)
	fmt.Println(res2)

	res3, err := SumAny(33.44, 55.66)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(res3)

	res4, err := SumAny("10", "20")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(res4) // 字符串拼接
}

func Sum(a, b int) int {
	return a + b
}

func SumFloat64(a, b float64) float64 {
	return a + b
}

func SumAny(a, b any) (any, error) {
	va, vb := reflect.ValueOf(a), reflect.ValueOf(b)
	if va.Kind() != vb.Kind() {
		return nil, errors.New("参数a和参数b类型不一致。")
	}

	// 使用类型断言和 switch 处理不同类型
	switch v := a.(type) {
	case int:
		return v + b.(int), nil
	case int8:
		return v + b.(int8), nil
	case int16:
		return v + b.(int16), nil
	case int32:
		return v + b.(int32), nil
	case int64:
		return v + b.(int64), nil
	case uint:
		return v + b.(uint), nil
	case uint8:
		return v + b.(uint8), nil
	case uint16:
		return v + b.(uint16), nil
	case uint32:
		return v + b.(uint32), nil
	case uint64:
		return v + b.(uint64), nil
	case float32:
		return v + b.(float32), nil
	case float64:
		return v + b.(float64), nil
	case string:
		return v + b.(string), nil
	default:
		return nil, errors.New("未知类型。")
	}
}
