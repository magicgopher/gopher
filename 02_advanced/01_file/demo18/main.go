package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 递归读取目录
	// 指定要读取的目录
	dirPath := "./testdir"

	// 递归遍历目录
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 如果出错，返回错误
		}

		// 打印路径和类型
		typeStr := "文件"
		if info.IsDir() {
			typeStr = "目录"
		}
		fmt.Printf("%s (%s)\n", path, typeStr)
		return nil
	})

	if err != nil {
		fmt.Println("遍历目录失败:", err)
		return
	}
}
