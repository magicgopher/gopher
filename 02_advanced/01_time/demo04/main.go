package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Location 类型
	// 表示与地区相关的时区，一个 time.Location 可能表示多个时区

	now := time.Now() // 获取本地当前时间

	// 获取本地时区
	// 方式一
	local1 := time.Local
	fmt.Println(local1.String(), now.In(local1)) // Local 2026-01-10 20:51:57.746277 +0800 CST

	// 方式二
	local2, _ := time.LoadLocation("Local")
	fmt.Println(local2.String(), now.In(local2))

	// 方式三
	local3 := now.Location()
	fmt.Println(local3.String(), now.In(local3))

	// 获取零时区
	// 方式一
	utc1 := time.UTC
	fmt.Println(utc1.String(), now.In(utc1))

	// 方式二
	utc2, _ := time.LoadLocation("")
	fmt.Println(utc2.String(), now.In(utc2))

	// 方式三
	utc3, _ := time.LoadLocation("UTC")
	fmt.Println(utc3.String(), now.In(utc3))

	shangahi, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(shangahi.String(), now.In(shangahi))

	beijing := time.FixedZone("Beijing", 8*60*60)
	fmt.Println(beijing.String(), now.In(beijing))
}
