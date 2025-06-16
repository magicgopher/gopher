package singleton

import (
	"fmt"
	"sync"
	"testing"
)

// TestGetInstance 测试单例
func TestGetInstance(t *testing.T) {
	// 获取单例实例
	instance1 := GetInstance()
	instance2 := GetInstance()
	// 判断实例是否相同
	if instance1 != instance2 {
		fmt.Println("instance1 is not equal to instance2")
	}
}

// TestConcurrentGetInstance 并发获取单例测试
func TestConcurrentGetInstance(t *testing.T) {
	// 定义并发执行的 goroutine 数量
	parCount := 100
	// 创建一个 sync.WaitGroup 用于等待所有 goroutine 完成
	wg := sync.WaitGroup{}
	// 创建一个存储单例实例的数组，用于保存每个 goroutine 获取的单例对象
	instances := make([]*Singleton, parCount)

	// 启动并发 goroutines 来获取单例实例
	for i := 0; i < parCount; i++ {
		wg.Add(1) // 每启动一个 goroutine，都增加一个计数
		go func(index int) {
			defer wg.Done()                  // 每个 goroutine 完成后，调用 Done() 来减少计数
			instances[index] = GetInstance() // 获取单例实例并存储在对应位置
		}(i)
	}

	// 等待所有的 goroutines 完成
	wg.Wait() // 阻塞，直到所有 goroutine 调用 Done() 完成

	// 检查所有 goroutine 中获取的单例实例是否相同
	for i := 1; i < parCount; i++ {
		// 如果某个实例不与第一个实例相同，说明单例实现存在问题
		if instances[i] != instances[0] {
			t.Errorf("Expected all instances to be the same, but instance %d is different.", i)
		}
	}
}
