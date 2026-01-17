package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// *File中Read()方法读取文件内容示例

func main() {
	// 打开文件
	file, err := os.Open("02_advanced/01_file/123.txt")
	if err != nil {
		log.Fatalf("文件打开失败: %v\n", err)
	}

	// defer关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", err)
		} else {
			log.Println("文件关闭成功!")
		}
	}()

	// 创建一个512字节的缓冲区，用于存放从文件中读取的内容
	// 长度和容量都是512，Read最多能读取512字节
	buf := make([]byte, 512)

	// 从文件中读取内容到 buf 中
	// n 表示实际读取到的字节数（可能小于512）
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		log.Printf("文件内容读取失败: %v\n", err)
	}

	// 输出实际读取到的字节数，以及读取到的字符串内容
	// 注意必须使用 buf[:n]，只取真正读到的部分，否则会输出一堆空字符（\x00）
	fmt.Printf("读取了 %d 个字节\n内容是: %s\n", n, string(buf[:n]))

}
