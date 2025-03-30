package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os 包的 Create 函数创建文件
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("创建文件失败, Error:", err)
		return
	}

	// 输出打印信息
	fmt.Println(file.Name(), "文件成功！")
}
