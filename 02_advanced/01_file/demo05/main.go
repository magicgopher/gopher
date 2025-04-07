package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件，指定文件的打开模式为只写和创建，权限为 0666
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败, Error:", err)
		return
	}

	// 关闭文件
	file.Close()

	// 要写入的数据
	data := []byte("Hello, World!")

	// 在文件的第3个字节位置写入数据
	offset := int64(3)
	n, err := file.WriteAt(data, offset)
	if err != nil {
		fmt.Println("写入文件失败, Error:", err)
		return
	}
	fmt.Printf("成功写入 %d 个字节\n", n)
}
