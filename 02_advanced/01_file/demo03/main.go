package main

import (
	"fmt"
	"log"
	"os"
)

// 读取小文件示例
// 推荐使用 os.ReadFile()

func main() {
	// 读取小文件推荐 os.ReadFile()
	fileData, err := os.ReadFile("02_advanced/01_file/123.txt")
	if err != nil {
		log.Printf("文件读取失败: %v\n", err)
	}
	fmt.Printf("文件共 %d 字节\n", len(fileData))
	fmt.Printf("文件内容: %v\n", string(fileData))
}
