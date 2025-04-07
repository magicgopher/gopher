package main

import (
	"fmt"
	"os"
)

func main() {
	// 以只读方式打开文件
	file1, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("打开文件失败, Error:", err)
		return
	}

	// 关闭文件
	defer file1.Close()

	// 输出打印信息
	fmt.Println(file1.Name(), "文件打开成功！")

	// 以读写方式打开文件
	file2, err := os.OpenFile("example.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("打开文件失败, Error:", err)
		return
	}

	// 关闭文件
	defer file2.Close()

	// 输出打印信息
	fmt.Println(file2.Name(), "文件打开成功！")
}
