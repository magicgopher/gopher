package main

import (
	"fmt"
	"log"
	"os"
)

// 创建单个文件示例

func main() {
	// 使用OpenFile()来创建文件
	// OpenFile(name string, flag int, perm FileMode) (*File, error)
	// name: 文件路径（可以是绝对路径/可以是相对路径）。
	// flag: 文件打开模式。
	// perm: 文件权限。常用0644 所有者有读写权限，其他人只有读权限 (rw-r--r--)。
	file, err := os.OpenFile("文档.txt", os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("文件创建失败: %v\n", err)
	}
	fmt.Printf("文件创建成功，文件名称: %v\n", file.Name())
}
