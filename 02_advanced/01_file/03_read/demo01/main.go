package main

import (
	"fmt"
	"log"
	"os"
)

// 小文件读取示例（文件大小 < 10MB）

func main() {
	// 只读的方式打开文件
	file, err := os.ReadFile("文档.txt")
	if err != nil {
		log.Fatalf("文件打开失败: %v\n", err)
	}
	// 输出文件大小和文件内容
	fmt.Printf("文件大小: %v\n", len(file))
	fmt.Printf("文件内容: %v\n", string(file))
}
