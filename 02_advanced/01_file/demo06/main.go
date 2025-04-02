package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件，并指定文件的模式为只写、创建、清空
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败, Error:", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 写入文件
	n, err := file.WriteString("file.WriteString()：Hello, Go!")
	if err != nil {
		fmt.Println("写入失败, Error:", err)
		return
	}
	fmt.Printf("成功写入 %d 个字节\n", n)
}
