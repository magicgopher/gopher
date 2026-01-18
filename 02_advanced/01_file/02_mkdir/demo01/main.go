package main

import (
	"fmt"
	"log"
	"os"
)

// 创建单个目录示例

func main() {
	// Mkdir(name string, perm FileMode) error
	// name: 文件路径（可以是绝对路径/可以是相对路径）。
	// perm: 文件权限。常用0644 所有者有读写权限，其他人只有读权限 (rw-r--r--)。
	err := os.Mkdir("test", 0666)
	if err != nil {
		log.Fatalf("目录创建失败: %v\n", err)
	}
	fmt.Println("目录创建成功！")
}
