package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开文件，并指定文件的模式为只写、创建、清空
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		// 如果打开文件时发生错误，打印错误信息并返回。
		fmt.Println("文件打开失败, Error:", err)
		return
	}

	// 使用 defer 确保在函数退出时关闭文件
	defer file.Close()

	// 创建一个缓冲写入器
	writer := bufio.NewWriter(file)

	// 向文件中写入字节数组 "Hello, Go!"，并获取写入的字节数和错误信息
	n, err := writer.Write([]byte("Hello, Go!"))
	if err != nil {
		// 如果写入时发生错误，打印错误信息并返回。
		fmt.Println("写入失败, Error:", err)
		return
	}

	// 刷新缓冲区
	writer.Flush()
	// 打印成功写入的字节数。
	fmt.Printf("成功写入 %d 个字节\n", n)
}
