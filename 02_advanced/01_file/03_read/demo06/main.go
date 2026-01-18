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
	// 使用 Reader Read() 按块读取
	fmt.Println("--- 开始按块读取 ---")
	// 创建一个 1KB 的桶
	buf := make([]byte, 1024)
	for {
		// 尝试读取数据填满 buf，n 是实际读到的字节数
		n, err := reader.Read(buf)
		// 处理读到的数据 (只处理前 n 个字节)
		if n > 0 {
			fmt.Printf("读到了 %d 个字节, 内容: %v\n", n, buf[:n])
			// 如果是文本，可以强转 string(buf[:n])
			//fmt.Printf("读到了 %d 个字节, 内容: %v\n", n, string(buf[:n]))
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("读取出错: %v", err)
			break
		}
	}
}
