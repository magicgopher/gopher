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

	// 输出打印信息
	fmt.Println(file.Name(), "文件打开成功！")
}
