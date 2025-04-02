package main

import (
	"fmt"
	"os"
)

func main() {
	// 这里不能使用 os.Open 函数打开文件，只有读取文件权限
	// 所以这里使用 os.OpenFile 函数打开文件，指定文件的开模式
	// os.O_RDONLY：只读｜ os.O_WRONLY：只写｜ os.O_RDWR：读写｜ os.O_CREATE：创建｜ os.O_APPEND：追加｜ os.O_TRUNC：清空
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败, Error:", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 写入文件
	n, err := file.Write([]byte("file.Write([]byte)：Hello, Go!"))
	if err != nil {
		fmt.Println("写入失败, Error:", err)
		return
	}
	fmt.Printf("成功写入 %d 个字节\n", n)
}
