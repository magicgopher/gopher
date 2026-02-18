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

// TestParseTime 将字符串格式时间解析为Time
func TestParseTime(t *testing.T) {
	timeStr1 := "2022-01-30 21:00:00"
	//timeStr1 := "Hello, World！"
	parse, err := time.Parse("2006-01-02 15:04:05", timeStr1)
	if err != nil {
		t.Fatalf("解析时间错误: %v\n", err)
	}
	t.Logf("解析后的时间: %v\n", parse)
}

// TestTimeFormat 将time转为字符串格式时间
func TestTimeFormat(t *testing.T) {
	date := time.Date(2020, 2, 2, 10, 15, 45, 0, time.Local)
	format := date.Format("2006年~01月~02日 15点04分05秒")
	t.Logf("类型: %T, 值: %v\n", format, format)
}

// TestAddTime 增加或者减少时间
func TestAddTime(t *testing.T) {
	date1 := time.Date(2026, 2, 1, 10, 0, 0, 0, time.Local)
	resultTime1 := date1.Add(-3 * time.Hour) // date1时间减三小时
	t.Logf("%v 3小时前的时间是: %v\n", date1, resultTime1)
	resultTime2 := date1.Add(5*time.Hour + 25*time.Minute)
	t.Logf("%v 5小时25分后的时间是: %v\n", date1, resultTime2)
}

// TestSubTime 两个时间的差值
func TestSubTime(t *testing.T) {
	t1 := time.Now()
	date1 := t1.Add(-3 * time.Hour) // 3小时前
	date2 := t1.Add(3 * time.Hour)  // 3小时后
	// “从过去到现在” → 现在.Sub(过去) → 正
	// “从现在到过去” → 过去.Sub(现在) → 负
	// “谁晚谁放前面，谁早谁放括号里”
	sub := date1.Sub(date2) // 从 date2 到 date1 过了多久？
	t.Log(sub)
}
