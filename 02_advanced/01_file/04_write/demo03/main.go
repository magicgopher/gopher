package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 打开文件，指定文件的打开模式为只写和创建，权限为 0666
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败, Error:", err)
		return
	}
	// defer关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", closeErr)
		} else {
			fmt.Println("文件关闭成功！")
		}
	}()
	// 要写入的数据
	data := []byte("Hello, World!")
	// 在文件的第3个字节位置写入数据
	offset := int64(3)
	n, err := file.WriteAt(data, offset)
	if err != nil {
		fmt.Println("写入文件失败, Error:", err)
		return
	}
	// 输出打印信息
	fmt.Printf("成功写入 %d 个字节\n", n)
}
