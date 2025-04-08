package main

import (
	"fmt"
	"os"
)

func main() {
	// 重命名文件
	err := os.Rename("example.txt", "example-new.txt")
	if err != nil {
		fmt.Println("文件重命名失败:", err)
		return
	}

	// 输出打印信息
	fmt.Println("文件重命名成功!")
}
