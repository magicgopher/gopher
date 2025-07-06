package singleton

import "sync"

// 使用 once 实现单例模式

var (
	instance *singleton // 单例结构体指针
	once     sync.Once  // 确保 instance 只初始化一次
)

// singleton 单例结构体
type singleton struct {
}

// GetInstance 返回 singleton 指针变量（单例实例）
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
