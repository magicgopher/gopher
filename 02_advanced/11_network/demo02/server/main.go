package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

// handleConn 处理客户端连接
func handleConn(conn net.Conn) {
	defer func() { // 关闭连接
		if closeErr := conn.Close(); closeErr != nil {
			log.Printf("关闭连接失败: %v\n", closeErr)
		} else {
			log.Println("关闭连接成功.")
		}
	}()
	remoteAddr := conn.RemoteAddr().String()
	log.Printf("新连接来自: %s", remoteAddr)
	// 使用 bufio 更方便按行读取
	reader := bufio.NewReader(conn)
	for {
		// 读取数据，直到遇到换行符 '\n'
		msg, err := reader.ReadString('\n')
		if err != nil {
			// 如果客户端主动断开连接，会返回 io.EOF 错误
			if err == io.EOF {
				log.Printf("客户端 [%s] 已断开连接\n", remoteAddr)
			} else {
				log.Printf("读取客户端数据失败: %v\n", err)
			}
			return // 发生错误或断开连接时，退出当前协程的循环
		}
		// 去掉读取到的字符串末尾的回车换行符
		msg = strings.TrimSpace(msg)
		log.Printf("收到来自客户端 [%s] 的消息: %s\n", remoteAddr, msg)
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
