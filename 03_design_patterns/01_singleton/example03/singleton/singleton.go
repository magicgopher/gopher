package singleton

import "sync"

// 创建实例时加速操作实现单例模式

var (
	instance *singleton // 单例结构体指针
	mutex    sync.Mutex // 互斥锁
)

// singleton 单例结构体
type singleton struct {
}

// GetInstance 返回 singleton 指针变量（单例实例）
func GetInstance() *singleton {
	if instance == nil { // 由于这里需要判断是否有实例，因此在多个 goroutine 获取实例时会有并发安全问题
		mutex.Lock()         // 加锁
		defer mutex.Unlock() // 最后释放锁
		instance = &singleton{}
	}
	return instance
}
