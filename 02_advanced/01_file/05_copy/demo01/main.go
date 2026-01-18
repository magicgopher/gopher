package main

import (
	"fmt"
	"log"
	"os"
)

// func (f *File) ReadFrom(r io.Reader) (n int64, err error)
// 用来从一个io.Reader源中读取数据，并直接写入到*File 中，它高效地将数据流从一个地方搬到另一个地方

func main() {
	// 1. 打开源文件（数据从哪来）
	srcFile, _ := os.Open("文档.txt")
	defer func() {
		if closeErr := srcFile.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", closeErr)
		}
	}()
	// 2. 创建目标文件（数据去哪儿）
	destFile, _ := os.Create("备份.txt")
	defer func() {
		if closeErr := destFile.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", closeErr)
		}
	}()
	// 3. 目标文件调用 ReadFrom，从源文件读取内容
	// 效果相当于把 srcFile 的内容“读”进了 destFile
	n, err := destFile.ReadFrom(srcFile) // io.Copy()底层也是
	if err != nil {
		log.Printf("拷贝过程中发生错误: %v\n", err)
	}
	fmt.Printf("成功拷贝 %d 字节\n", n)
}
