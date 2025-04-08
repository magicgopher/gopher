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
		fmt.Println("文件打开失败, Error:", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 创建一个字节切片作为读取的缓冲区，缓冲区大小为15
	buf := make([]byte, 15)

	// 从第5个字节开始读取
	n, err := file.ReadAt(buf, 5)
	if err != nil && err != io.EOF {
		fmt.Println("文件读取失败, Error:", err)
		return
	}

	// 打印实际读取的字节数
	fmt.Printf("从偏移量5读取了 %d 个字节.\n", n)

	// 将读取到的字节转换为字符串并打印
	fmt.Printf("读取到的内容:%s\n", string(buf[:n]))
}
