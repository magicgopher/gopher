package main

import (
	"fmt"
	"os"
)

func main() {
	// 删除文件
	err := os.Remove("example.txt")
	if err != nil {
		fmt.Println("文件删除失败:", err)
		return
	}

	// 输出打印信息
	fmt.Println("文件删除成功!")
}
