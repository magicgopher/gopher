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
	fmt.Println("--- 开始按行读取 ---")
	for {
		// 读取直到遇到换行符 '\n'
		line, err := reader.ReadString('\n')
		// 注意：line 里面通常包含了结尾的 '\n'，打印时一般不需要额外换行
		fmt.Println(line)
		// 处理文件结束 (EOF)
		if err != nil {
			if err == io.EOF {
				break // 文件读完了，退出循环
			}
			log.Printf("读取出错: %v\n", err)
			break
		}
	}
}
