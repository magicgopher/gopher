package strconvt

import (
	"strconv"
	"testing"
)

// TestStringToInteger 字符串转为int类型
func TestStringToInteger(t *testing.T) {
	s1 := "12345"
	got, err := strconv.Atoi(s1)
	if err != nil {
		t.Errorf("strconv.Atoi(%q) 返回了错误: %v", s1, err)
		return
	}
	// 期望值
	want := 12345
	if got != want {
		t.Errorf("strconv.Atoi(%q) = %d, 期望 %d", s1, got, want)
	}
	// 打印结果
	t.Logf("成功转换: %q → %d (类型: %T)", s1, got, got)
}
