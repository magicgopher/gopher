package example // 这是声明在example包下

import "fmt"

// SayHello 打招呼函数
// name表示字符串类型的参数
func SayHello(name string) {
	// 这里使用fmt包的Printf函数格式化控制台输出打印
	// %s表示字符串占位符
	fmt.Printf("Hello, %s!\n", name)
}
