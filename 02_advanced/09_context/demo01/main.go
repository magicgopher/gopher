package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 取消信号
// 场景：同时向 Google 和 Baidu 发 http 请求，整体超时 3 秒，哪个先超时或被取消就立刻停止那个请求。

func main() {
	// 创建一个带有 3 秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // 无论如何 3 秒后都会调用 cancel
	var wg sync.WaitGroup
	wg.Add(2)
	// 第一个 goroutine：请求 Google
	go func() {
		defer wg.Done()
		err := fetchWithContext(ctx, "https://www.google.com", "Google")
		if err != nil {
			fmt.Println("Google 请求失败或被取消:", err)
		}
	}()
	// 第二个 goroutine：请求 Baidu
	go func() {
		defer wg.Done()
		err := fetchWithContext(ctx, "https://www.baidu.com", "Baidu")
		if err != nil {
			fmt.Println("Baidu 请求失败或被取消:", err)
		}
	}()
	// 等待两个 goroutine 结束，或者整体超时
	wg.Wait()
	fmt.Println("主程序结束")
}

// fetchWithContext 带有 context 控制的 http 请求
func fetchWithContext(ctx context.Context, url, name string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{
		Timeout: 10 * time.Second, // 单次请求最长10秒（但会被 ctx 更早打断）
	}
	fmt.Printf("[%s] 开始请求...\n", name)
	resp, err := client.Do(req)
	if err != nil {
		// 这里 err 很可能是 context.DeadlineExceeded 或 context.Canceled
		return fmt.Errorf("%s 请求出错: %w", name, err)
	}
	defer resp.Body.Close()
	// 模拟只读一点点 body（实际场景可以全部读完）
	_, err = io.ReadAll(io.LimitReader(resp.Body, 1024*32)) // 最多读 32KB
	if err != nil {
		return fmt.Errorf("%s 读取响应体失败: %w", name, err)
	}
	fmt.Printf("[%s] 请求成功，状态码: %d\n", name, resp.StatusCode)
	return nil
}
