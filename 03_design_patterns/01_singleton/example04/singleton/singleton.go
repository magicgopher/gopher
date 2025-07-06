package singleton

import "sync"

// 双重检查加锁实现单例模式

var (
	instance *singleton // 单例结构体指针
	mutex    sync.Mutex // 互斥锁
)

// singleton 单例结构体
type singleton struct {
}

// GetInstance 返回 singleton 指针变量（单例实例）
func GetInstance() *singleton {
	if instance == nil { // 第一次检查没有加锁，减少锁的开销
		mutex.Lock()         // 加锁
		defer mutex.Unlock() // 最后释放锁
		if instance == nil { // 第二次检查，确保只有一个 goroutine 可以创建实例
			instance = &singleton{}
		}
	}
	return instance
}
