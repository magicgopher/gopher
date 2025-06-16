package singleton

import "sync"

// 单例模式

var instance *Singleton
var once sync.Once

// Singleton singleton结构体
type Singleton struct {
}

// GetInstance 获取Singleton指针实例
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
