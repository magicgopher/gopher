package main

import (
	"fmt"
	"log"
	"os"
)

// 获取打开的文件信息示例

func main() {
	file, err := os.Open("文档.txt")
	if err != nil {
		log.Fatalf("文件打开失败: %v\n", err)
	}
	// 关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", closeErr)
		} else {
			fmt.Println("文件关闭成功！")
		}
	}()
	// 获取文件信息
	fileInfo, err := file.Stat()
	fmt.Println("文件名:", fileInfo.Name())
	fmt.Println("文件大小:", fileInfo.Size())
	fmt.Println("文件权限:", fileInfo.Mode())
	fmt.Println("文件修改时间:", fileInfo.ModTime())
	fmt.Println("文件是否是目录:", fileInfo.IsDir())
	fmt.Println("文件是否是文件:", fileInfo.Mode().IsRegular())
}
