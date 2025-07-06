package singleton

var instance *singleton // 单例结构体指针

// singleton 单例结构体
type singleton struct {
}

// init 函数
func init() {
	instance = &singleton{} // 初始化单例
}

// GetInstance 返回 singleton 指针变量（单例实例）
func GetInstance() *singleton {
	return instance
}
