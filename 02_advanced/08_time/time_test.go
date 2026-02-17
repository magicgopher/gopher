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

// TestGetYearMonthDay 获取年月日
func TestGetYearMonthDay(t *testing.T) {
	now := time.Now()    // 获取本地时间，这里会返回一个 Time
	year := now.Year()   // 获取年份
	month := now.Month() // 获取月份
	day := now.Day()     // 获取日期
	t.Logf("%v 年 %v 月 %v 日\n", year, month, day)
}

// TestGetHourMinuteSecond 获取时分秒
func TestGetHourMinuteSecond(t *testing.T) {
	now := time.Now()      // 获取本地时间，这里会返回一个 Time
	hour := now.Hour()     // 获取当前小时
	minute := now.Minute() // 获取分钟
	second := now.Second() // 获取秒
	t.Logf("%v 时 %v 分 %v 秒\n", hour, minute, second)
}

// TestGetTime 创建一个特定的Time实例
func TestGetTime(t *testing.T) {
	date := time.Date(2020, 2, 2, 10, 15, 45, 0, time.Local)
	t.Log(date)
}
