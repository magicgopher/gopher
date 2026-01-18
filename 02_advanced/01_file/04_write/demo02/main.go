package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 打开文件，指定文件的打开模式为只写和创建，权限为 0666
	file, err := os.OpenFile("文档.txt", os.O_WRONLY|os.O_CREATE, 0666)
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
	// os.File 的 WriteString 方法用于将字符串写入文件
	n, err := file.WriteString("Hello, World!")
	if err != nil {
		fmt.Println("字符串内容写入失败, Error:", err)
		return
	}
	// 输出打印信息
	fmt.Println("文件写入成功！写入了", n, "个字节")
}
