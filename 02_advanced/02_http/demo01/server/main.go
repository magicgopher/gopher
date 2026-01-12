package main

import (
	"fmt"
	"log"
	"net/http"
)

// 这是 server 端

func Hello(w http.ResponseWriter, r *http.Request) {
	// 只允许 GET 请求
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 处理查询参数 Query Parameters
	param := r.URL.Query().Get("param")

	// 如果没有提供 param 参数，就返回默认值
	if param == "" {
		param = "World!"
	}

	// 拼接响应内容
	resp := fmt.Sprintf("Hello, %s!", param)

	// 写入响应
	n, err := w.Write([]byte(resp))
	if err != nil {
		log.Printf("写入响应失败: %v\n", err)
		return
	}
	log.Printf("成功响应客户端，param=%s，写入字节数=%d\n", param, n)
}

func main() {
	// 配置路由，当访问 /hello 时调用 Hello 函数处理请求
	http.HandleFunc("/hello", Hello)

	log.Println("服务器启动，监听 :8080")
	// 启动服务端并监听8080端口
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("服务器启动失败: %v\n", err)
	}
}
