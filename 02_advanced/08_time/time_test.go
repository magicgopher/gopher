package _8_time

import (
	"testing"
	"time"
)

// TestTimeNow 获取当前系统时间（当前本地时间）
func TestTimeNow(t *testing.T) {
	now := time.Now() // 获取当前时间
	t.Log(now)
}

// TestGetWeekday 获取星期几
func TestGetWeekday(t *testing.T) {
	now := time.Now() // 获取当前时间
	t.Log(now.Weekday())
}

// TestGetYear 获取年月日
func TestGetYear(t *testing.T) {
	now := time.Now()    // 获取本地时间，这里会返回一个 Time
	year := now.Year()   // 获取年份
	month := now.Month() // 获取月份
	day := now.Day()     // 获取日期
	t.Logf("%v 年 %v 月 %v 日", year, month, day)
}
