package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取 Time 基本信息

	now := time.Now()

	// 地点和时区信息
	fmt.Println(now.Location()) // Local

	// 时区
	fmt.Println(now.Zone()) // CST 28800
	// Unix 时间戳（单位：秒）
	fmt.Println(now.Unix())

	// 获取当前是星期几
	weekday := now.Weekday()
	fmt.Println(weekday)

	// 返回当前 time.Time 对应的 ISO 8601 周编号
	year, w := now.ISOWeek()
	fmt.Println(year, w) // year表示ISO周所在的年份，w是这一年的第几周

	date := time.Date(2020, 8, 20, 19, 30, 45, 0, now.Location())
	// Clock() 返回当前时间的时分秒
	fmt.Println(date.Clock()) // 19 30 45

	// YearDay() 返回当前日期是这一年的第几天（1-based，从1月1日开始算第1天）
	// 范围 1-365或者366
	yearDay := now.YearDay()
	fmt.Println(yearDay)
}
