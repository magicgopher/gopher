package stringst

import (
	"strings"
	"testing"
)

// TestStringClone 复制字符串
func TestCloneString(t *testing.T) {
	src := "hello, world!"
	t.Logf("原来字符串的地址: %p, 字符串内容: %v\n", &src, src)
	result := strings.Clone(src)
	t.Logf("原来字符串的地址: %p, 字符串内容: %v\n", &result, result)
}

// TestCompareString 比较字符串
func TestCompareString(t *testing.T) {
	s1 := "abc"
	s2 := "abc"
	s3 := "abe"
	s4 := "ab"
	f := func(s string) int {
		sum := 0
		for _, r := range s {
			sum += int(r)
		}
		return sum
	}
	i1 := f(s1)
	i2 := f(s2)
	t.Logf("s1码值: %v, s2码值: %v\n", i1, i2)
	t.Log(strings.Compare(s1, s2)) // 字符串1 == 字符串2（码值总和相等），结果: 0
	i3 := f(s3)
	t.Logf("s1码值: %v, s3码值: %v\n", i1, i3)
	t.Log(strings.Compare(s1, s3)) // 字符串1 < 字符串2 结果: -1
	i4 := f(s4)
	t.Logf("s1码值: %v, s4码值: %v\n", i1, i4)
	t.Log(strings.Compare(s1, s4)) // 字符串1 > 字符串2 结果: 1
}
