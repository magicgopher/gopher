package main

import (
	"fmt"
	"os"
)

func main() {
	// 指定要读取文件夹的路径
	dirPath := "./testdir"

	// 读取目录内容
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("读取目录失败:", err)
		return
	}

	// 遍历目录条目
	fmt.Println("目录内容：")
	for i, entry := range entries {
		// 获取条目名称和类型
		name := entry.Name()
		isDir := entry.IsDir()
		fmt.Printf("%d. %s (%s)\n", i+1, name, map[bool]string{true: "目录", false: "文件"}[isDir])
	}
}
