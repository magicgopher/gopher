package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 使用 os.OpenFile 以只读模式打开文件
	// os.O_RDONLY: 只读
	// 0644: 文件权限（在只读模式下，权限位通常不生效，但必须提供）
	file, err := os.OpenFile("文档.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	// defer关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", closeErr)
		} else {
			fmt.Println("文件关闭成功！")
		}
	}()
	// 定义一个缓冲区，大小为5，用于存放读取文件的数据
	buf := make([]byte, 5)
	offset := int64(0)
	// 循环读取
	for {
		n, err := file.ReadAt(buf, offset)
		// n大于0证明还有数据
		if n > 0 {
			fmt.Printf("读取到 %d 字节数据, 内容是: %v\n", n, string(buf[:n]))
			offset += int64(n) // 重新设置偏移量，原始的偏移量+读取的字节
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
