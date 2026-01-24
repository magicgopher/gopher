package main

import (
	"log"
	"net/http"
)

// 使用 go 内置 http 标准库启动一个web服务

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	// req.Method 获取请求的方法
	// req.URL.Path 获取请求的路径
	log.Printf("请求的方法: %v, 处理请求路径为: %v\n", req.Method, req.URL.Path)

	// 响应请求
	write, err := resp.Write([]byte("Hello World!"))
	if err != nil {
		log.Printf("写入响应失败: %v\n", err)
		return
	}
	log.Printf("写入响应长度: %d\n", write)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("启动服务器并监听 :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
