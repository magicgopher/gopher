package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Printf("关闭连接失败: %v\n", closeErr)
		} else {
			log.Println("连接已关闭.")
		}
	}()
	inputReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)
	for {
		fmt.Print("请输入 (半双工): ")
		input, _ := inputReader.ReadString('\n')
		// 1. 客户端先发送数据
		_, err := conn.Write([]byte(input))
		if err != nil {
			log.Printf("向服务端发送数据失败: %v", err)
			break
		}
		// 2. 阻塞等待服务端的回复 (这就是半双工的体现，收发不能同时进行)
		reply, err := serverReader.ReadString('\n')
		if err != nil {
			log.Println("读取服务端回复失败")
			break
		}
		fmt.Printf(">> 收到回复: %s", reply)
	}
}
