package main

import (
	"context"
	"fmt"
	"time"
)

// 场景：调用一个可能很慢的外部服务，最多允许等 4 秒，超时就全部放弃。

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	// 模拟一个很慢的操作（实际可能是数据库、RPC、第三方 API 等）
	go func() {
		time.Sleep(6 * time.Second) // 故意超过 4 秒
		fmt.Println("慢操作居然完成了（但你看不到，因为已经超时取消了）")
	}()

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("主程序等了5秒")
	case <-ctx.Done():
		fmt.Println("超时了！原因:", ctx.Err()) // 输出: context deadline exceeded
	}
}
