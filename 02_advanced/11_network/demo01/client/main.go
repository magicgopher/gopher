package main

import (
	"fmt"
	"log"
	"net"
)

// 客户端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("连接服务器失败: %v", err)
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Printf("关闭连接失败: %v", closeErr)
		} else {
			log.Println("连接正常关闭")
		}
	}()
	// 客户端发送一句话
	message := "你好，服务端，这是数据 5050\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println("发送失败:", err)
		return
	}
	fmt.Printf("已发送: %v\n", message)
	// 接收服务端回复（只读一行）
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("读取回复失败:", err)
		return
	}
	reply := string(buf[:n])
	fmt.Println("收到服务端回复:", reply)
}
