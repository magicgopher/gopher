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

	// os.File 的 Write 方法用于将字节数组写入文件
	n, err := file.Write([]byte("Hello, World!"))
	if err != nil {
		fmt.Println("字节数组内容写入失败, Error:", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 输出打印信息
	fmt.Println("文件写入成功！写入了", n, "个字节")
}
