package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// 半双工：数据可以在两个方向上传输，但不能同时进行。同一时刻，要么是 A 发给 B，要么是 B 发给 A
// 生活例子：对讲机（按下按钮说话，松开按钮听对方说，不能同时抢话）

// handleConn 处理客户端连接
func handleConn(conn net.Conn) {
	defer func() { // 关闭连接
		if closeErr := conn.Close(); closeErr != nil {
			log.Printf("关闭连接失败: %v\n", closeErr)
		} else {
			log.Println("关闭连接成功.")
		}
	}()
	// 缓冲读，用于获取客户端发送过来的数据
	reader := bufio.NewReader(conn)
	for {
		// 等待读取客户端的消息
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("客户端断开连接")
			}
			return
		}
		msg = strings.TrimSpace(msg)
		log.Printf("收到客户端消息: %s", msg)
		// 处理完毕后，给客户端发送回复
		reply := fmt.Sprintf("服务端已收到你的消息 [%s]\n", msg)
		_, err = conn.Write([]byte(reply))
		if err != nil {
			log.Printf("回复客户端失败: %v", err)
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("服务端启动失败: %v\n", err)
	}
	defer func() {
		if closeErr := listen.Close(); closeErr != nil {
			log.Printf("服务端关闭失败: %v\n", closeErr)
		}
	}()
	log.Println("服务端启动成功, 端口: 8888")
	// 循环监听
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Printf("接收连接失败: %v\n", err)
			continue
		}
		// 开启新的 goroutine 处理该连接
		go handleConn(accept)
	}
}
