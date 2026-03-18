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
		log.Fatalf("无法连接到服务端: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("关闭连接失败: %v", err)
		}
	}()
	// 开启独立的 Goroutine 负责【接收服务端的主动推送】
	go func() {
		serverReader := bufio.NewReader(conn)
		for {
			msg, err := serverReader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					log.Println("\n服务端正常关闭了连接")
				} else {
					log.Printf("\n读取服务端推送异常: %v\n", err)
				}
				// 服务端断开了，客户端也没有继续运行的意义了
				os.Exit(0)
			}
			// 为了不打乱正在输入的提示符，先换行打印消息，再补上输入提示
			fmt.Print("\n" + msg + "请输入 (全双工): ")
		}
	}()
	// 主协程负责【读取键盘输入并发送】
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入 (全双工): ")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("输入结束，准备退出")
			} else {
				log.Printf("读取键盘输入失败: %v", err)
			}
			break
		}
		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Printf("发送数据到服务端失败: %v", err)
			break
		}
	}
}
