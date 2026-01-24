package main

import (
	"fmt"
	"log"
	"net/http"
)

// 服务端
// 处理客户端的请求已经响应数据

// Hello 接收客户端查询参数，例如：/xxx?value=xxx
func Hello(w http.ResponseWriter, r *http.Request) {
	// 简单的CORS处理
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 只接收GET请求
	if r.Method != "GET" {
		http.Error(w, "请求方法有误！", http.StatusMethodNotAllowed)
	}
	// 解析参数
	value := r.URL.Query().Get("value")
	// value参数为空，那就给默认值
	if value == "" {
		value = "this is the server!"
	}
	// 拼接响应字符内容
	content := fmt.Sprintf("Hi, %s!", value)
	// 响应请求
	write, err := w.Write([]byte(content))
	if err != nil {
		log.Printf("响应写入内容失败: %v\n", err)
		return
	}
	fmt.Printf("成功写入响应 %d 字节\n", write)
}

func main() {
	// 定义路由接口
	http.HandleFunc("/hello", Hello)
	log.Println("服务端启动成功")
	// 启动服务端并监听8080端口
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("服务端启动失败: %v", err)
	}
}
