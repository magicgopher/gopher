package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 读取文件内容示例
// 使用 *File 的Read()方法读取文件内容

func main() {
	// 以只读的方式打开文件
	file, err := os.Open("文档.txt")
	if err != nil {
		log.Fatalf("文件打开失败: %v\n", err)
	}
	// defer关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", closeErr)
		} else {
			fmt.Println("文件关闭成功！")
		}
	}()
	// 创建字节切片作为读取的缓冲区，缓冲区大小为128
	buf := make([]byte, 128)
	for {
		// 读取文件，这里就像是用水勺也就是buf去水桶里面勺水
		n, err := file.Read(buf)
		// 先处理读到的数据（只要 n > 0）就正面还有数据
		if n > 0 {
			// 输出读取字节大小
			fmt.Printf("文件读取成功！读取了 %d 个字节\n", n)
			// 输出读取到的文件内容
			fmt.Printf("读取到的内容: %v\n", string(buf[:n]))
		}
		// 错误处理
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件已经全部读取完了。")
				break
			}
			log.Printf("文件读取失败: %v\n", err)
			break
		}
	}
}
