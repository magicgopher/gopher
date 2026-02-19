package strconvt

import (
	"log"
	"strconv"
	"testing"
)

// TestStringToInt 整数类型字符串转为int类型
func TestStringToInt(t *testing.T) {
	s1 := "12345"
	//s1 := "123.45" // 浮点数字符串
	result, err := strconv.Atoi(s1)
	if err != nil {
		log.Printf("转换失败: %v\n", err)
		return
	}
	t.Logf("类型: %T, 值: %v\n", result, result)
}

// TestIntegerToString int类型转整数类型字符串
func TestIntegerToString(t *testing.T) {
	result := strconv.Itoa(78910)
	t.Logf("类型: %T, 值: %v\n", result, result)
}

// TestStringToBool 字符串转bool类型
func TestStringToBool(t *testing.T) {
	// "1", "t", "T", "true", "TRUE", "True" // true
	// "0", "f", "F", "false", "FALSE", "False" // false
	result, err := strconv.ParseBool("true")
	//result, err := strconv.ParseBool("F")
	if err != nil {
		log.Printf("转换失败: %v\n", err)
		return
	}
	t.Logf("类型: %T, 值: %v\n", result, result)
}
