package main

import (
	"fmt"
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

	// 创建一个字节切片作为读取的缓冲区，缓冲区大小为128
	buf := make([]byte, 128)
	// 调用 Read 方法读取文件内容到缓冲区
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("文件读取失败, Error:", err)
		return
	}

	// 打印实际读取的字节数
	fmt.Println("文件读取成功！读取了", n, "个字节")

	// 将读取到的字节转换为字符串并打印
	fmt.Println("读取到的内容:", string(buf[:n]))
}
