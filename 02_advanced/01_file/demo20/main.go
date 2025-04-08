package main

import (
	"fmt"
	"os"
)

func main() {
	// 指定要删除的目录
	dirPath := "testdir"

	// 创建测试目录和内容
	err := os.MkdirAll(dirPath+"/subdir", 0755)
	if err != nil {
		fmt.Println("创建目录失败:", err)
		return
	}
	err = os.WriteFile(dirPath+"/file1.txt", []byte("Hello"), 0644)
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}

	// 递归删除目录及其内容
	err = os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println("删除目录失败:", err)
		return
	}

	fmt.Println("目录及其内容删除成功！")
}
