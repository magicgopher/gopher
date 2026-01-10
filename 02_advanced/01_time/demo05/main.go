package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Time 类型
	// 表示时间，可以通过下列函数获取time.Time对象，然后通过该对象获取年月日时分秒等信息。

	// Now()：获取本地当前时间
	now := time.Now()
	fmt.Println(now)

	// Date：自定义时间
	date := time.Date(2016, 6, 1, 12, 0, 0, 0, now.Location())
	fmt.Println(date)
}
