package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// 全双工：数据可以在两个方向上同时传输。双方都可以随时发送和接收数据
// 生活例子：打电话（双方可以同时说话和听到对方的声音）

// handleConn 处理连接
func handleConn(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("关闭连接失败: %v", err)
		}
	}()
	remoteAddr := conn.RemoteAddr().String()
	// 开启独立的 Goroutine 负责【主动推送】
	go func() {
		for {
			time.Sleep(5 * time.Second)
			msg := "【服务端广播】当前时间是: " + time.Now().Format("15:04:05") + "\n"
			_, err := conn.Write([]byte(msg))
			if err != nil {
				// 当客户端断开连接时，这里的 Write 会报错（broken pipe 或 connection reset）
				log.Printf("向客户端 [%s] 推送数据失败，停止推送协程: %v", remoteAddr, err)
				return // 停止这个 goroutine，防止泄露
			}
		}
	}()
	// 主协程负责【读取接收】
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("客户端 [%s] 正常断开", remoteAddr)
			} else {
				log.Printf("读取客户端 [%s] 数据异常: %v", remoteAddr, err)
			}
			return
		}
		log.Printf("收到客户端 [%s] 消息: %s", remoteAddr, strings.TrimSpace(msg))
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("全双工服务端启动失败: %v", err)
	}
	defer func() {
		if err := listen.Close(); err != nil {
			log.Printf("关闭监听器失败: %v", err)
		}
	}()
	log.Println("全双工服务端启动, 端口: 8888")
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Printf("接收连接失败: %v", err)
			continue
		}
		go handleConn(accept)
	}
}
