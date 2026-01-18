package main

import (
	"fmt"
	"log"
	"os"
)

// 打开文件/关闭文件示例

func main() {
	// Open(name string) (*File, error)
	// 以只读的模式打开文件
	// name: 文件路径（可以是绝对路径/可以是相对路径）。
	file, err := os.Open("文档.txt")
	if err != nil {
		log.Fatalf("文件打开失败: %v\n", err)
	}
	// defer关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", closeErr)
		} else {
			fmt.Println("文件关闭成功！")
		}
	}()
}
