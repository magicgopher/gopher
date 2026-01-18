package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 使用OpenFile()打开/创建文件
	// os.O_WRONLY: 只写
	// os.O_CREATE: 文件不存在就创建
	// os.O_TRUNC: 文件若存在，清空它重新写
	file, err := os.OpenFile("缓冲写测试文档.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败:", err)
		return
	}
	// defer关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", closeErr)
		} else {
			fmt.Println("文件关闭成功！")
		}
	}()
	// 创建缓冲 Writer
	writer := bufio.NewWriter(file)
	n, err := writer.WriteString("Hello, World! (这是通过 WriteString 写入的)\n")
	if err != nil {
		fmt.Println("写入出错:", err)
	}
	fmt.Printf("成功写入 %d 字节\n", n)
}
