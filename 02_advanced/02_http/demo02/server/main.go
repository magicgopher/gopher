package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response 响应体格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// UserReq 用户查询请求
type UserReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// User 用户模型（模拟数据库中的用户模型）
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func SendJSON(w http.ResponseWriter, code int, message string, data interface{}) {
	// 设置响应头，告诉客户端返回的是 JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // HTTP 状态码通常都返回 200，具体错误看业务 code

	resp := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	// 将结构体编码为 JSON 并写入 ResponseWriter
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("JSON编码失败: %v", err)
	}
}

func UserSearch(w http.ResponseWriter, r *http.Request) {
	// 只允许Post请求
	if r.Method != http.MethodPost {
		SendJSON(w, 405, "Method Not Allowed", nil)
		return
	}
	// 解析请求体 JSON
	var req UserReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendJSON(w, 400, "请求参数格式错误", nil)
		return
	}

	log.Printf("收到查询请求: Name=%s, Age=%d", req.Name, req.Age)

	// 模拟业务逻辑（查库）
	mockUser := User{
		ID:   1001,
		Name: "张三锋",
		Age:  22,
		Sex:  "男",
	}

	// 返回统一响应
	SendJSON(w, 200, "查询成功", mockUser)
}

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 方法限制：只允许 GET
	if r.Method != http.MethodGet {
		SendJSON(w, 405, "Method Not Allowed", nil)
		return
	}

	// 2. 模拟业务逻辑（查库获取列表）
	userList := []User{
		{ID: 1, Name: "小红", Age: 30, Sex: "女"},
		{ID: 2, Name: "小黑", Age: 32, Sex: "男"},
		{ID: 3, Name: "小蓝", Age: 28, Sex: "男"},
		{ID: 4, Name: "小橙", Age: 28, Sex: "男"},
		{ID: 5, Name: "小美", Age: 28, Sex: "女"},
	}

	// 3. 返回统一响应
	SendJSON(w, 200, "获取列表成功", userList)
}

func main() {
	// 注册查询接口
	http.Handle("/api/v1/user/", http.HandlerFunc(UserSearch))

	// 注册查询列表接口
	http.Handle("/api/v1/user/list", http.HandlerFunc(UserListHandler))

	log.Println("服务器启动，监听 :8080")
	// 启动服务器并监听8080端口
	if listenErr := http.ListenAndServe(":8080", nil); listenErr != nil {
		log.Printf("服务器启动失败: %v\n", listenErr)
	}
}
