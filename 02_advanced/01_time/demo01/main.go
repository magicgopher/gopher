package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()          // 获取当前时间
	utcNow := time.Now().UTC() // UTC时间
	fmt.Println(now)
	fmt.Println(utcNow)
}
