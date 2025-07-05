package example02

import (
	"sync"
	"testing"
)

// TestSingleGoroutine 测试单个goroutine获取单例实例
func TestSingleGoroutine(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 != instance2 {
		t.Fatal("instance1不等于instance2")
	}
}

// TestMultipleGoroutines 测试多个goroutine并发获取单例实例
func TestMultipleGoroutines(t *testing.T) {
	// 重置instance以确保测试的独立性
	instance = nil

	// 定义goroutine数量
	const goroutineCount = 1000

	// 用于同步goroutine
	var wg sync.WaitGroup
	wg.Add(goroutineCount)

	// 存储每个goroutine获取的实例
	instances := make([]*singleton, goroutineCount)

	// 启动多个goroutine并发获取单例实例
	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()
			instances[index] = GetInstance()
		}(i)
	}

	// 等待所有goroutine执行完成
	wg.Wait()

	// 检查所有实例是否相同
	for i := 1; i < goroutineCount; i++ {
		if instances[i] != instances[0] {
			t.Logf("goroutine %d 获取的实例地址: %p, 第一个实例地址: %p", i, instances[i], instances[0])
			t.Fatalf("goroutine %d 获取的实例与第一个实例不同，存在并发安全问题", i)
		}
	}
}
