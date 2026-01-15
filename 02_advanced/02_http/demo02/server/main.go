package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// H1 处理 Query Parameters (?name=xxx)
func H1(w http.ResponseWriter, r *http.Request) {
	// 只允许GET请求
	if r.Method != "GET" {
		http.Error(w, "请求方法有误", http.StatusMethodNotAllowed)
		return
	}

	// 解析URL上的参数（Query Parameters）
	query := r.URL.Query()

	// 获取具体的参数值
	name := query.Get("name")
	age := query.Get("age")
	address := query.Get("address")
	log.Printf("收到请求: name=%s, age=%s, address=%s\n", name, age, address)

	// 构建一个响应体数据
	respData := map[string]interface{}{
		"code": 200,
		"msg":  "获取成功",
		"data": map[string]string{
			"name":    name,
			"age":     age,
			"address": address,
		},
	}

	// 设置响应头 (Content-Type)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// 返回JSON数据
	if err := json.NewEncoder(w).Encode(respData); err != nil {
		log.Printf("JSON 编码失败: %v", err)
	}
}

// H2 处理 Path Parameters (/world/{id}/{name}...)
func H2(w http.ResponseWriter, r *http.Request) {
	// 只允许GET请求
	if r.Method != "GET" {
		http.Error(w, "请求方法有误", http.StatusMethodNotAllowed)
		return
	}

	// 获取URL的路径参数
	id := r.PathValue("id")
	name := r.PathValue("name")
	age := r.PathValue("age")
	address := r.PathValue("address")
	log.Printf("H2 收到请求: id=%s, name=%s, age=%s, address=%s\n", id, name, age, address)

	// 构建响应
	respData := map[string]interface{}{
		"code": 200,
		"msg":  "H2 Path参数获取成功",
		"data": map[string]string{
			"id":      id,
			"name":    name,
			"age":     age,
			"address": address,
		},
	}

	// 设置响应头 (Content-Type)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// 返回JSON数据
	if err := json.NewEncoder(w).Encode(respData); err != nil {
		log.Printf("JSON 编码失败: %v", err)
	}
}

func main() {
	// 示例：GET请求 http://localhost:8080/api/v1/hello?name="张三"&age=20&address="广东广州"
	http.Handle("/api/v1/hello", http.HandlerFunc(H1))

	// 示例：GET请求 http://localhost:8080/api/v1/world/1/张三/20/广东广州
	http.Handle("/api/v1/world/{id}/{name}/{age}/{address}", http.HandlerFunc(H2))

	log.Println("服务端启动，并监听 8080 端口")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("服务端启动失败: %v\n", err)
	}
}
