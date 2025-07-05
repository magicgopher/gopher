package example02

import "sync"

// 单例模式
// 加互斥锁的方式实现

var (
	instance *singleton // 单例结构体指针全局变量
	mutex    sync.Mutex // 互斥锁
)

// singleton 单例结构体
type singleton struct {
}

// GetInstance 返回 singleton 的唯一实例
func GetInstance() *singleton {
	mutex.Lock()         // 函数开始执行就加锁，保证并发的安全性，但是性能会有所下降
	defer mutex.Unlock() // 最后释放锁
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
