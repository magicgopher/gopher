package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败, Error:", err)
		return
	}
	defer file.Close()

	n, err := file.WriteAt([]byte("Go"), 7) // 从第7字节开始写入
	if err != nil {
		fmt.Println("写入失败, Error:", err)
		return
	}
	fmt.Printf("成功写入 %d 个字节\n", n)
}
