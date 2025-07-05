package singleton

import "time"

// 单例模式
// 不加锁实现

// 定义全局变量 singleton指针变量
var instance *singleton

// singleton 单例结构体
type singleton struct {
}

// GetInstance 返回 *singleton 指针变量（单例实例）
func GetInstance() *singleton {
	if instance == nil { // 当实例为nil时创建单例实例，这里在多个goroutine获取实例时会有并发安全问题
		time.Sleep(800 * time.Millisecond)
		instance = &singleton{}
	}
	return instance
}
