package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	fmt.Println("文件名:", fileInfo.Name())
	fmt.Println("文件大小:", fileInfo.Size())
	fmt.Println("文件权限:", fileInfo.Mode())
	fmt.Println("文件修改时间:", fileInfo.ModTime())
	fmt.Println("文件是否是目录:", fileInfo.IsDir())
	fmt.Println("文件是否是文件:", fileInfo.Mode().IsRegular())
}
