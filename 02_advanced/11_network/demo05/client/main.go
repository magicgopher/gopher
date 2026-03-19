package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("连接服务端失败: %v\n", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("关闭连接失败: %v\n", err)
		}
	}()
	// 开启独立的 Goroutine 专门负责【接收服务端的消息】(包括系统提示和别人的私信)
	go func() {
		serverReader := bufio.NewReader(conn)
		for {
			msg, err := serverReader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					log.Println("\n与服务端断开连接")
				} else {
					log.Printf("\n读取服务端消息异常: %v\n", err)
				}
				os.Exit(0)
			}
			// 打印收到的消息，并重新打印输入提示符
			fmt.Print(msg + "请输入: ")
		}
	}()
	// 主协程专门负责【读取键盘输入并发送】
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 第一个输入的内容会被服务端当做“用户名”注册
		// 后续输入的内容需要遵循 "目标用户:消息内容" 的格式
		input, err := inputReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("准备退出...")
			} else {
				log.Printf("读取键盘输入失败: %v", err)
			}
			break
		}
		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Printf("发送数据失败: %v", err)
			break
		}
	}
}
