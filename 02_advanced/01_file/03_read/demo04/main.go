package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 使用 os.OpenFile 以只读模式打开文件
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
	// 定义读取文件存储的缓存区大小为20
	buf := make([]byte, 20)
	// 先读取前几个字节
	n, _ := file.Read(buf)
	fmt.Printf("开头读取: %d 字节 → %q\n", n, buf[:n])
	// 跳回文件最开头再读一次
	file.Seek(0, io.SeekStart)
	n, _ = file.Read(buf)
	fmt.Printf("跳回开头再读: %d 字节 → %q\n", n, buf[:n])
	// 跳到文件末尾
	file.Seek(0, io.SeekEnd)
	pos, _ := file.Seek(0, io.SeekCurrent) // 查看当前位置
	fmt.Printf("当前在文件末尾，第 %d 字节处\n", pos)
	// 从末尾往前移动 8 个字节（读最后8个字）
	file.Seek(-8, io.SeekEnd)
	n, _ = file.Read(buf)
	fmt.Printf("文件最后8字节: %q\n", buf[:n])
	// 读取剩余所有内容（常用模式）
	file.Seek(0, io.SeekStart) // 先回开头
	content, _ := io.ReadAll(file)
	fmt.Printf("完整内容:\n%s\n", content)
}
