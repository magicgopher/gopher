package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // 获取当前时间

	year := now.Year()           // 年
	month := now.Month()         // 月 这里会显示英文
	monthInt := int(now.Month()) // 月 显示 int类型月份
	day := now.Day()             // 日
	hour := now.Hour()           // 时
	minute := now.Minute()       // 分
	second := now.Second()       // 秒
	fmt.Printf("%v年-%v月-%v日 %v时:%v分:%v秒\n", year, month, day, hour, minute, second)
	fmt.Printf("%v年-%v月-%v日 %v时:%v分:%v秒\n", year, monthInt, day, hour, minute, second)
}
