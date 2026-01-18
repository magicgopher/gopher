package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 只读的方式打开文件
	file, err := os.Open("文档.txt")
	if err != nil {
		log.Fatalf("文件打开失败: %v\n", err)
	}
	// 关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", closeErr)
		} else {
			fmt.Println("文件关闭成功！")
		}
	}()
	// 创建一个带缓冲的 Reader，默认缓冲区大小通常是 4KB
	reader := bufio.NewReader(file)
	// 使用 ReadByte 逐字节读取
	fmt.Println("--- 开始逐字节读取 ---")
	for {
		// 读取 1 个字节
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("读取出错: %v\n", err)
			break
		}
		// 打印字符及其 ASCII 码
		fmt.Printf("字符: %s (ASCII: %d)\n", string(b), b)
	}
}
