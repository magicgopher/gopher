package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开文件，指定文件的打开模式为只写和创建，权限为 0666
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE, 066)
	if err != nil {
		fmt.Println("文件打开失败, Error:", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 创建一个带有默认缓冲区大小的 bufio.Writer
	writer := bufio.NewWriter(file)

	// 要写入的数据
	data := []string{
		"这是第一行内容。\n",
		"这是第二行内容。\n",
		"这是第三行内容。\n",
		"这是第四行内容。\n",
		"这是第五行内容。\n",
	}

	// 使用 Write 方法将数据写入缓冲区
	for _, line := range data {
		n, err := writer.Write([]byte(line))
		if err != nil {
			fmt.Println("写入到缓冲区时发生错误, Error:", err)
			return
		}
		fmt.Printf("写入了 %d 个字节到缓冲区, 内容是: %q\n", n, line)
	}

	// 重要：调用 Flush 方法将缓冲区中的数据写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区到文件时发生错误, Error:", err)
		return
	}
	fmt.Println("数据已成功写入到文件!")
}
