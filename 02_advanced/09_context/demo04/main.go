package main

import (
	"context"
	"fmt"
)

// 传递请求范围的值(数据透传)
// 场景：在 HTTP 处理链路中，把 traceID、userID 等透传到所有下游函数。

type traceKey string

const traceIDKey traceKey = "traceID"

func main() {
	ctx := context.Background()
	// 在入口处设置值（比如 middleware 里）
	context.WithValue(ctx, traceIDKey, "req-abc123-xyz789")
	// 层层调用
	handleRequest(ctx)
}

// handleRequest
func handleRequest(ctx context.Context) {
	// 中间业务逻辑
	fmt.Println("处理请求中...")
	// 往下传
	callDatabase(ctx)
	callExternalAPI(ctx)
}

// callDatabase
func callDatabase(ctx context.Context) {
	if v := ctx.Value(traceIDKey); v != nil {
		if traceID, ok := v.(string); ok {
			fmt.Printf("[DB] 使用 traceID: %s 执行查询\n", traceID)
			return
		}
	}
	fmt.Println("[DB] 警告：未找到 traceID，使用默认或生成新的")
}

// callExternalAPI
func callExternalAPI(ctx context.Context) {
	// 最常用、最简洁的安全写法
	if traceID, ok := ctx.Value(traceIDKey).(string); ok && traceID != "" {
		fmt.Printf("[API] 携带 traceID: %s 调用第三方\n", traceID)
	} else {
		fmt.Println("[API] 警告：未找到 traceID")
		// 可以在这里记录日志、生成随机 traceID 等
	}
}
