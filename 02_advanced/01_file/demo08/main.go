package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os.WriteFile 函数将字节数组写入文件
	// 参数：
	//   - "example.txt": 文件名
	//   - []byte("Hello, Go!"): 要写入的字节数组，这里是将字符串 "Hello, Go!" 转换为字节数组
	//   - 0666: 文件权限，0666 表示所有用户都具有读写权限
	err := os.WriteFile("example.txt", []byte("Hello, Go!"), 0666)

	// 检查写入过程中是否发生错误
	if err != nil {
		// 如果发生错误，则打印错误信息并返回
		fmt.Println("写入失败, Error:", err)
		return
	}

	// 如果没有发生错误，则打印文件写入成功的消息
	fmt.Println("文件写入成功！")
}
