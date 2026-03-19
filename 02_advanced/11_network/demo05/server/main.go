package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

// 服务端

// clientMap 作为全局通讯录，保存 "用户名" 到 "TCP 连接" 的映射
var clientMap = make(map[string]net.Conn)

// mapMutex 用于保护 clientMap 的并发读写安全
var mapMutex sync.RWMutex

// handleConn 处理客户端连接
func handleConn(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("关闭连接异常: %v\n", err)
		}
	}()
	// 创建缓冲读
	reader := bufio.NewReader(conn)
	// 要求客户端连上后，第一句话必须是自己的用户名
	_, err := conn.Write([]byte("已连接！请输入你的用户名进行注册:\n"))
	if err != nil {
		return
	}
	nameLine, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("读取用户名失败: %v\n", err)
		return
	}
	name := strings.TrimSpace(nameLine)
	// 将新用户加入通讯录
	mapMutex.Lock()
	if _, exists := clientMap[name]; exists {
		conn.Write([]byte("该用户名已存在，连接断开。\n"))
		mapMutex.Unlock()
		return
	}
	clientMap[name] = conn
	mapMutex.Unlock()
	log.Printf("用户上线: [%s], IP: %s\n", name, conn.RemoteAddr().String())
	conn.Write([]byte(fmt.Sprintf("注册成功！欢迎 [%s]。\n(发送消息格式: 目标用户名:你的消息)\n", name)))
	// 当函数退出时（即客户端断开连接），从通讯录中移除该用户
	defer func() {
		mapMutex.Lock()
		delete(clientMap, name)
		mapMutex.Unlock()
		log.Printf("用户下线: [%s]\n", name)
	}()
	// 消息转发阶段：循环读取该客户端发来的消息
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("客户端 [%s] 主动断开\n", name)
			} else {
				log.Printf("读取 [%s] 数据异常: %v\n", name, err)
			}
			return
		}
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}
		// 解析协议，格式要求："目标用户:消息内容" (使用中文或英文冒号皆可，这里以英文冒号为例)
		parts := strings.SplitN(msg, ":", 2)
		if len(parts) < 2 {
			conn.Write([]byte("发送失败！格式错误，请使用 '目标用户名:消息内容'\n"))
			continue
		}
		targetName := strings.TrimSpace(parts[0])
		content := strings.TrimSpace(parts[1])
		// 去通讯录中查找目标用户
		mapMutex.RLock() // 只读锁
		targetConn, ok := clientMap[targetName]
		mapMutex.RUnlock()
		if ok {
			// 找到了目标用户，将消息转发过去
			forwardMsg := fmt.Sprintf("\n【来自 %s 的私信】: %s\n", name, content)
			_, err := targetConn.Write([]byte(forwardMsg))
			if err != nil {
				log.Printf("转发给 [%s] 失败: %v\n", targetName, err)
				conn.Write([]byte(fmt.Sprintf("发送给 [%s] 失败，对方网络异常\n", targetName)))
			} else {
				// 给发送者一个成功反馈
				conn.Write([]byte(fmt.Sprintf("-> 已成功发送给 [%s]\n", targetName)))
			}
		} else {
			// 没找到目标用户
			conn.Write([]byte(fmt.Sprintf("发送失败：用户 [%s] 不在线或不存在\n", targetName)))
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
			log.Printf("关闭服务端失败: %v\n", closeErr)
		}
	}()
	log.Println("服务端启动成功，端口: 8888，等待客户端连接...")
	//
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Printf("接收客户端连接失败: %v\n", err)
			continue
		}
		// 启动goroutine处理连接
		go handleConn(accept)
	}
}
