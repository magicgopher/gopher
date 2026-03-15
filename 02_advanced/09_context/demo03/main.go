package main

import (
	"context"
	"fmt"
	"time"
)

// 截止时间
// 场景：这个任务必须在“今天下午 3 点前”完成，过了就没意义了。

func main() {
	// 假设现在是 14:58，我们设置截止到 15:02
	deadline := time.Now().Add(4 * time.Minute) // 4 分钟后截止
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	fmt.Println("截止时间是:", deadline.Format("15:04:05"))
	// 模拟长时间任务
	for i := 0; i < 10; i++ {
		select {
		case <-time.After(30 * time.Second):
			fmt.Println("完成一步...", i+1)
		case <-ctx.Done():
			fmt.Println("超过截止时间了！原因:", ctx.Err())
			// 这里通常会清理资源、记录日志等
			return
		}
	}
	fmt.Println("全部完成（在截止时间内）")
}
