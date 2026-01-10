package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()        // 获取当前时间
	unix := now.Unix()       // 秒级时间戳
	milli := now.UnixMilli() // 毫秒级时间戳
	micro := now.UnixMicro() // 微秒级时间戳
	nano := now.UnixNano()   //纳秒级时间戳
	fmt.Printf("秒级时间戳:%v\n", unix)
	fmt.Printf("毫秒级时间戳:%v\n", milli)
	fmt.Printf("微秒级时间戳:%v\n", micro)
	fmt.Printf("纳秒级时间戳:%v\n", nano)
}
