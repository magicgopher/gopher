package main

import (
	"log"
	"os"
)

// 打开文件和关闭文件示例

func main() {
	// 打开文件
	file, err := os.Open("02_advanced/01_file/123.txt")
	if err != nil {
		log.Fatalf("文件打开失败, err:%v\n", err)
	}
	log.Println("成功打开文件.")

	// defer 关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("文件关闭失败, err:%v\n", closeErr)
		} else {
			log.Printf("文件成功关闭.")
		}
	}()
}
