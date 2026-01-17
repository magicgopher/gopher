package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// ReadAt(b []byte, off int64) (n int, err error)
// 从文件的指定偏移量位置读取数据，不改变文件指针位置

func main() {
	// 打开文件
	file, err := os.OpenFile("02_advanced/01_file/123.txt", os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("文件打开失败: %v\n", err)
	}

	// defer关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("文件关闭失败: %v\n", closeErr)
		} else {
			log.Println("文件关闭成功！")
		}
	}()

	// 创建一个字节切片作为读取的缓冲区，缓冲区大小为10
	buf := make([]byte, 5)
	offset := int64(0)
	for {
		n, err := file.ReadAt(buf, offset)

		// n大于0就还有数据，可以一直读取
		if n > 0 {
			// 每次读取输出信息
			fmt.Printf("从偏移量 %d 开始读取了 %d 个字节.\n", offset, n)
			fmt.Printf("读取到的内容:%s\n", string(buf[:n]))
			// 修改偏移量
			offset += int64(n)
		}

		// 文件读取完毕
		if errors.Is(err, io.EOF) {
			fmt.Println("文件读取完毕 (EOF)")
			break
		}

		// 文件读取失败
		if err != nil {
			log.Fatalf("文件读取失败: %v\n", err)
		}

	}
}
