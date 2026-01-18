package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// 模拟多行文本文件
	fileContent := "第一行内容\n第二行内容\n第三行内容"
	// 1. 创建 Scanner (注意这里是 NewScanner，不是 NewReader)
	scanner := bufio.NewScanner(strings.NewReader(fileContent))
	fmt.Println("--- 开始使用 Scanner 读取 ---")
	// 2. 这里的循环非常简洁：只要还有下一行，Scan() 就返回 true
	for scanner.Scan() {
		// 3. 获取当前行的内容
		// scanner.Text() 返回字符串
		// scanner.Bytes() 返回字节切片
		text := scanner.Text()
		fmt.Println("读到一行:", text)
	}
	// 4. 循环结束后，一定要检查一下是否有非 EOF 的错误
	if err := scanner.Err(); err != nil {
		fmt.Printf("读取过程中发生了错误: %v\n", err)
	}
	fmt.Println("--- 读取结束 ---")
}
