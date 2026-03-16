package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 服务端

// handleConn 处理单个连接
func handleConn(conn net.Conn) {
	// 确保连接一定会被关闭
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Printf("关闭连接失败 %s → %s : %v\n",
				conn.RemoteAddr().String(), conn.LocalAddr().String(), closeErr)
		} else {
			log.Printf("连接正常关闭 %s → %s\n",
				conn.RemoteAddr().String(), conn.LocalAddr().String())
		}
	}()
	// 读取客户端发来的内容（只读一行）
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println("读取失败:", err)
		return
	}
	// 去掉末尾的换行符，方便看
	msg := line
	if len(msg) > 0 && msg[len(msg)-1] == '\n' {
		msg = msg[:len(msg)-1]
	}
	fmt.Println("收到客户端:", msg)
	// 简单处理：从消息里提取数字部分（这里我们假装知道格式）
	// 实际项目中可以用 strings.Split 或更严格的解析
	reply := "服务端收到数据 5050\n"
	// 发回给客户端
	_, err = conn.Write([]byte(reply))
	if err != nil {
		log.Println("写回失败:", err)
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
	// 循环监听客户端连接
	for {
		listener, err := listen.Accept()
		if err != nil {
			log.Printf("服务端接收连接失败: %v\n", err)
			continue
		}
		// 每来一个连接就处理
		handleConn(listener)
	}
}
