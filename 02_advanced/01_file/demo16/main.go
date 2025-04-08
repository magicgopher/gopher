package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件夹
	err := os.Mkdir("a", 0755)
	if err != nil {
		fmt.Println("文件夹创建失败:", err)
		return
	}
	fmt.Println("文件夹创建成功！")

	// 创建多层级文件夹
	err = os.MkdirAll("./a/b/c", 0755)
	if err != nil {
		fmt.Println("多层级文件夹创建失败:", err)
		return
	}
	fmt.Println("多层级文件夹创建成功！")
}
