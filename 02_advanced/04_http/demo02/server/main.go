package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 服务端

func H1(w http.ResponseWriter, r *http.Request) {
	// 只能处理GET请求
	if r.Method != "GET" {
		http.Error(w, "只支持 GET 请求！", http.StatusMethodNotAllowed)
	}
	// 解析GET请求的查询参数(Query Parameters)
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	address := r.URL.Query().Get("address")
	log.Printf("name=%s, age=%s, address=%s", name, age, address)
	// 将参数拼接成响应信息
	msg := fmt.Sprintf("我是%s, 年龄: %s, 地址:%s", name, age, address)
	// 定义一个统一响应的结构体
	resp := struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: http.StatusOK,
		Msg:  "请求成功！",
		Data: msg,
	}
	// 设置响应头
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 序列化为json并写入响应
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "服务器内部错误", http.StatusInternalServerError)
		return
	}
}

func H2(w http.ResponseWriter, r *http.Request) {
	// 只能处理GET请求
	if r.Method != http.MethodGet {
		http.Error(w, "只支持 GET 请求！", http.StatusMethodNotAllowed)
		return
	}
	// 获取路径参数
	name := r.PathValue("name")
	age := r.PathValue("age")
	address := r.PathValue("address")
	log.Printf("路径参数: name=%s, age=%s, address=%s", name, age, address)
	// 将参数拼接成响应信息（与 H1 保持一致）
	msg := fmt.Sprintf("我是%s, 年龄: %s, 地址:%s", name, age, address)
	// 定义一个统一响应的结构体
	resp := struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: http.StatusOK,
		Msg:  "请求成功！",
		Data: msg,
	}
	// 设置响应头
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 序列化为json并写入响应
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "服务器内部错误", http.StatusInternalServerError)
		return
	}
}

func main() {
	// 路由定义
	// 处理GET请求 查询参数(query parameters) http://127.0.0.1:8080/api/v1/user?name=xxx=age=xx=address=xxx
	http.Handle("/api/v1/user", http.HandlerFunc(H1))
	// 处理GET请求 路径参数(path parameters) http://127.0.0.1:8080/api/v1/user/{name}/{age}/{address}
	http.Handle("/api/v1/user/{name}/{age}/{address}", http.HandlerFunc(H2))
	log.Println("服务端启动成功，并监听8080端口。")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("服务端启动失败: %v\n", err)
	}
}
