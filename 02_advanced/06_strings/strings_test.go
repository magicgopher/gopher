package stringst

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
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
	t.Logf("字符串 s 中 空字符串出现 %v 次\n", strings.Count(s, ""))
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
	// 中文字符通常占用 3 个字节、emoji 通常占 4 字节
	s := "hello 你好 world! 😊 123"
	// Index: 查找完整的一段字符串（子串）
	t.Log(strings.Index(s, "你好"))     // 6
	t.Log(strings.Index(s, "你好 123")) // 没有找到，结果: -1
	t.Log(strings.Index(s, "l"))      // 2
	t.Log("==================")
	// IndexAny: 找一组字符中的任意一个
	t.Log(strings.IndexAny(s, "你好"))      // 6
	t.Log(strings.IndexAny(s, "好 world")) // 2
	t.Log(strings.IndexAny(s, "你"))       // 6
	t.Log("==================")
	// IndexRune：精确找单个 rune
	t.Log(strings.IndexRune(s, '你')) // 6     （'你' 的起始字节位置）
	t.Log(strings.IndexRune(s, '好')) // 9     （'好' 的字节位置）
	t.Log(strings.IndexRune(s, '😊')) // 20    （emoji 通常占 4 字节）
	t.Log(strings.IndexRune(s, 'z')) // -1
}

// TestMap 遍历替换字符串
func TestMap(t *testing.T) {
	s1 := "abc"
	// 大小写替换
	result1 := strings.Map(func(r rune) rune {
		return r - 32
	}, s1)
	t.Log(result1)
	// 删除某些字符串
	s2 := "Order #12345 - Total: $99.99"
	result2 := strings.Map(func(r rune) rune {
		// 判断是否是数字字符
		if r >= '0' && r <= '9' {
			return -1
		}
		return r
	}, s2)
	t.Log(result2)
	// 替换特定的字符
	s3 := "123$456#789%987#654$321"
	result3 := strings.Map(func(r rune) rune {
		if r == '$' || r == '#' || r == '%' {
			r = '-'
		}
		return r
	}, s3)
	t.Log(result3)
	//
	s4 := "Hello, 世界! 2026 Go😊"
	result4 := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == ' ' {
			return r
		}
		return -1
	}, s4)
	t.Log(result4)
	// 处理中文、emoji 等 Unicode 字符
	s5 := "你好，世界！😊"
	result5 := strings.Map(func(r rune) rune {
		if r == '你' {
			return '我'
		}
		if r == '😊' {
			return '👍'
		}
		return r
	}, s5)
	t.Log(result5)
}

// TestRepeatString 重复拼接一个字符串
func TestRepeatString(t *testing.T) {
	result1 := strings.Repeat("abc", 3) // 将abc字符串复制三次，然后拼接再一起
	t.Log(result1)
	result2 := strings.Repeat("你好", 2) // 将你好字符复制两次，然后拼接再一起
	t.Log(result2)
}

// TestReplaceString 替换字符串
func TestReplaceString(t *testing.T) {
	// n：表示替换次数，-1表示全部替换，0表示不替换
	s1 := "Hello, Java."
	replace1 := strings.Replace(s1, "Java", "Golang", 1)
	t.Log(replace1)
	s2 := "C++ C++ C++"
	replace2 := strings.Replace(s2, "C++", "Python", 2)
	t.Log(replace2)
	replace3 := strings.Replace(s2, "C++", "Python", -1)
	t.Log(replace3)
	replace4 := strings.Replace(s2, "C++", "Python", 0)
	t.Log(replace4)
	replace5 := strings.ReplaceAll(s2, "C++", "Golang")
	t.Log(replace5)
}

// TestSplitString 分割字符串
func TestSplitString(t *testing.T) {
	s1 := "123|456|789|987|654|321"
	result1 := strings.Split(s1, "|")
	t.Log(result1)
	s2 := "Java$Go$Python$C$Rust"
	// SplitN函数的n参数表示分割的次数
	// n>0 表示最多返回 n 个子串，第 n 个子串将包含剩余所有未分割的内容。
	// n<0 表示返回所有可能的子串，没有数量限制。
	// n=0 表示返回nil切片
	result2 := strings.SplitN(s2, "$", 2)
	t.Log(result2)
	s3 := "a,b,c,d,e,f"
	result3 := strings.SplitAfter(s3, ",")
	t.Log(result3)
	result4 := strings.SplitAfterN(s3, ",", 3)
	t.Log(result4)
	// SplitAfterSeq 函数是1.24引入的新函数
	// 旨在利用 Go 新加入的迭代器 (Iterator) 机制来高效地分割字符串，分割的字符串包含其后的分隔符
	for s := range strings.SplitAfterSeq(s3, ",") {
		t.Log(s)
	}
	// SplitSeq 和 SplitAfterSeq 类似，只是SplitSeq分割后的子串不会保留分割符
	for s := range strings.SplitSeq(s3, ",") {
		t.Log(s)
	}
}

// TestToLowerAndUpper 大小写转换
func TestToLowerAndUpper(t *testing.T) {
	t.Log(strings.ToLower("My Name is MagicGopher!"))
	t.Log(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
	t.Log(strings.ToUpper("My name is jack,Nice to meet you!"))
	t.Log(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}

// TestTrimString 修剪字符串
func TestTrimString(t *testing.T) {
	s1 := "!!Hello, World!!!"
	result1 := strings.Trim(s1, "!")
	t.Log(result1)
	result2 := strings.TrimLeft(s1, "!")
	t.Log(result2)
	result3 := strings.TrimRight(s1, "!")
	t.Log(result3)
	result4 := strings.TrimFunc(s1, func(r rune) bool {
		return r == '!'
	})
	t.Log(result4)
	result5 := strings.TrimPrefix(s1, "!")
	t.Log(result5)
	result6 := strings.TrimSuffix(s1, "!")
	t.Log(result6)
	s2 := "  你好，Golang！  "
	result7 := strings.TrimSpace(s2)
	t.Log(result7)
}
