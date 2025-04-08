package main

import (
	"fmt"
	"os"
)

func main() {
	// 指定要删除的空目录
	dirPath := "emptydir"

	// 创建一个空目录用于测试
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		fmt.Println("创建目录失败:", err)
		return
	}

	// 删除空目录
	err = os.Remove(dirPath)
	if err != nil {
		fmt.Println("删除目录失败:", err)
		return
	}

	fmt.Println("空目录删除成功！")
}
