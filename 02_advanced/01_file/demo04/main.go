package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os 包的 Open 函数打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("文件打开失败, Error:", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败, Error:", err)
		return
	}

	// 打印文件信息
	fmt.Println("文件名:", fileInfo.Name())
	fmt.Println("文件大小:", fileInfo.Size())
	fmt.Println("修改时间:", fileInfo.ModTime())
	fmt.Println("文件权限:", fileInfo.Mode())
}
