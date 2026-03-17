package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// 连接到服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("连接服务端失败: %v\n", err)
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Printf("关闭连接失败: %v\n", closeErr)
		} else {
			log.Println("连接已关闭.")
		}
	}()
	log.Println("成功连接到服务端！可以直接输入内容 (输入 'exit' 退出程序):")
	// 创建一个从控制台（标准输入）读取内容的 Reader
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入发送的消息: ")
		// 读取用户输入，直到按下回车键（产生换行符）
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Printf("读取控制台输入失败: %v\n", err)
			continue
		}
		// 如果用户输入了 exit，就退出循环，结束客户端
		if strings.TrimSpace(input) == "exit" {
			log.Println("主动退出...")
			break
		}
		// 将读取到的数据发送给服务端
		// 注意：input 包含换行符，正好符合服务端按行读取 (ReadString('\n')) 的逻辑
		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Printf("发送数据失败: %v\n", err)
			break
		}
	}
}
