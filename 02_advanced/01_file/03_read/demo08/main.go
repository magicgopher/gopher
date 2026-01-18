package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	// 模拟文件内容：用逗号分隔的数据
	fileContent := "apple,banana,orange,grape"
	// 创建 Reader
	reader := bufio.NewReader(strings.NewReader(fileContent))
	fmt.Println("--- 开始使用 ReadBytes 读取 ---")
	for {
		// 读取直到遇到逗号 ','
		// 注意：返回的 bytes 中是包含逗号的！
		bytes, err := reader.ReadBytes(',')
		if len(bytes) > 0 {
			// 为了演示，我把它转成 string 打印出来
			// 实际开发中，你通常直接操作 bytes
			fmt.Printf("读到的字节数据: %v (转文字: %q)\n", bytes, string(bytes))
		}
		if err != nil {
			if err == io.EOF {
				fmt.Println("--- 文件读取完毕 ---")
				break
			}
			fmt.Printf("发生错误: %v\n", err)
			break
		}
	}
}
