package hello

import "fmt"

// go可见性原则是通过开头大小写决定的
// 大写表示可导出 “公有” 可以被其他包访问
// 小写表示不可导出 “私有” 只能在当前包内使用

type Hello struct {
	value string // 小写的外部访问不了
	Value string // 大写的，外部才可以访问
}

// SayHello hello函数
func SayHello(v string) string {
	return fmt.Sprintf("Hello, %v", v)
}

// sayHello hello函数
func sayHello(v string) string {
	return fmt.Sprintf("Hello, %v", v)
}

func (h *Hello) Hello() string {
	return h.value
}

func (h *Hello) hello() string {
	return h.value
}
