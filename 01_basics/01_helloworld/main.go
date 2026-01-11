package main // 包名，用于声明当前go文件属于哪个包下

import "fmt" // import是导入关键字，用于声明导入哪个包，这里导入了fmt包

// main函数是整个go程序的入口
func main() {
	// 调用了fmt包下的Println函数进行控制台输出
	fmt.Println("Hello World!")
}
