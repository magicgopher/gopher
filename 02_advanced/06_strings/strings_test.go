package stringst

import (
	"fmt"
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

// TestContainsString 包含字符串
func TestContainsString(t *testing.T) {
	s1 := "hdiosjdosajdoi"
	result1 := strings.Contains(s1, "sa")
	t.Logf("s1字符串是否包含: %s 字符串, 结果: %v\n", "sa", result1)
	result2 := strings.Contains(s1, "hello")
	t.Logf("s1字符串是否包含: %s 字符串, 结果: %v\n", "hello", result2)
	result3 := strings.Contains(s1, "sjdo")
	t.Logf("s1字符串是否包含: %s 字符串, 结果: %v\n", "sjdo", result3)
}

// TestCountSubString 子串出现次数
func TestCountSubString(t *testing.T) {
	s := "263817491469317401"
	t.Logf("字符串 s 中 1出现 %v 次\n", strings.Count(s, "1"))
	t.Logf("字符串 s 中 2出现 %v 次\n", strings.Count(s, "2"))
	t.Logf("字符串 s 中 99出现 %v 次\n", strings.Count(s, "99"))
}

// TestCutString 删除指定子字符串
func TestCutString(t *testing.T) {
	// 例子1：正常切分
	before, after, found := strings.Cut("hello,world", ",")
	fmt.Println(before, after, found) // 输出: hello world true
	// 例子2：切空格
	before, after, found = strings.Cut("Go is great", " ")
	fmt.Println(before, after, found) // 输出: Go is great true
	// 例子3：没找到分隔符
	before, after, found = strings.Cut("no-separator-here", ":")
	fmt.Println(before, after, found) // 输出: no-separator-here "" false
	// 例子4：分隔符在开头
	before, after, found = strings.Cut(":start", ":")
	fmt.Println(before, after, found) // 输出: "" start true
	// 例子5：分隔符在结尾
	before, after, found = strings.Cut("end:", ":")
	fmt.Println(before, after, found) // 输出: end "" true
	// 例子6：多次出现，只切第一个
	before, after, found = strings.Cut("a-b-c-d", "-")
	fmt.Println(before, after, found) // 输出: a b-c-d true
}

// TestEqualFold 判断两个字符串内容（忽略大小写）
func TestEqualFold(t *testing.T) {
	s1 := "hello"
	s2 := "hello"
	fold1 := strings.EqualFold(s1, s2)
	t.Logf("%s 字符串和 %s 字符串是否相等, 结果: %v\n", s1, s2, fold1)
	s3 := "Hello"
	fold2 := strings.EqualFold(s1, s3)
	t.Logf("%s 字符串和 %s 字符串是否相等, 结果: %v\n", s1, s3, fold2)
	s4 := "world"
	fold3 := strings.EqualFold(s1, s4)
	t.Logf("%s 字符串和 %s 字符串是否相等, 结果: %v\n", s1, s4, fold3)
}

// TestField 字符串分割
func TestField(t *testing.T) {
	t.Logf("%q\n", strings.Fields(" a b c d e f g "))
	t.Logf("%q\n", strings.FieldsFunc("a,b,c,d,e,f,g", func(r rune) bool {
		return r == ','
	}))
}

// TestPreSuffix 寻找前、后缀
func TestPreSuffix(t *testing.T) {
	str := "abbc cbba"
	fmt.Println(strings.HasPrefix(str, "abb"))
	fmt.Println(strings.HasSuffix(str, "bba"))
}

// TestIndexString 查找子串的位置
func TestIndexString(t *testing.T) {
	s1 := "abcdefg"
	index := strings.Index(s1, "d")
	fmt.Printf("字符串 %s 第一次出现的位置: %d\n", "d", index)
	s2 := "Hello, 世界! PHP是世界最好的语言。"
	index1 := strings.Index(s2, "d")
	index2 := strings.Index(s2, "世界")
	fmt.Printf("字符串 %s 第一次出现的位置: %d\n", "d", index1)  // 没有找到返回-1
	fmt.Printf("字符串 %s 第一次出现的位置: %d\n", "世界", index2) // 没有找到返回-1
}
