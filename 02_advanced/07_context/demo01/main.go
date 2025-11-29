package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// 超时控制的context简单示例

func main() {
	// 给整个程序最多 3 秒时间
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // 非常重要！

	fmt.Println("主程序启动，等待长任务...")
	err := longTask(ctx)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("任务超时了！", err)
		} else {
			fmt.Println("任务被取消了！", err)
		}
	} else {
		fmt.Println("任务正常完成")
	}
}

func longTask(ctx context.Context) error {
	fmt.Println("长任务开始执行...")
	select {
	case <-time.After(8 * time.Second): // 模拟耗时 8 秒的任务
		fmt.Println("长任务完成！")
		return nil
	case <-ctx.Done(): // 被取消或超时
		fmt.Println("长任务被中断")
		return ctx.Err()
	}
}
