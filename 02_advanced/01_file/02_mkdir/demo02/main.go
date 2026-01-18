package main

import (
	"fmt"
	"os"
)

// 级联(递归)创建多个目录示例

func main() {
	// 定义一个深层嵌套的路径
	dirPath := "tmp/upload/images"
	// 使用 MkdirAll 递归创建
	// path: 文件路径。
	// perm: 文件权限。常用0644 所有者有读写权限，其他人只有读权限 (rw-r--r--)。
	// 0755 权限：所有者(rwx) 组用户(rx) 其他用户(rx)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Printf("创建目录失败: %s\n", err)
		return
	}
	fmt.Printf("目录 %s 创建成功 (或者已存在)\n", dirPath)
}
