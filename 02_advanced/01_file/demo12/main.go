package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 将文件指针移动到第5个字节（从开头算起）
	_, err = file.Seek(5, io.SeekStart)
	if err != nil {
		fmt.Println("Seek 失败:", err)
		return
	}

	// 读取5个字节的内容
	buffer := make([]byte, 5)
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println("读取失败:", err)
		return
	}

	fmt.Printf("读取到的内容: %s\n", buffer[:n])
}
