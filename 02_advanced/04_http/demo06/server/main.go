package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// 服务端示例代码
// 处理POST请求（application/json）格式的数据

var users = []User{
	{ID: 1, Name: "张三", Age: 20, Sex: "男"},
	{ID: 2, Name: "张无忌", Age: 22, Sex: "男"},
	{ID: 3, Name: "张三丰", Age: 30, Sex: "男"},
	{ID: 4, Name: "李四", Age: 21, Sex: "男"},
	{ID: 5, Name: "王五", Age: 23, Sex: "男"},
	{ID: 6, Name: "赵六", Age: 19, Sex: "女"},
	{ID: 7, Name: "张丽", Age: 20, Sex: "女"},
	{ID: 8, Name: "孙七", Age: 25, Sex: "男"},
	{ID: 9, Name: "周八", Age: 24, Sex: "女"},
	{ID: 10, Name: "吴九", Age: 18, Sex: "男"},
}

// User 用户结构体
type User struct {
	ID   int64  `json:"id"`   // ID
	Name string `json:"name"` // 姓名
	Age  int    `json:"age"`  // 年龄
	Sex  string `json:"sex"`  // 性别
}

// UserQueryReq 用户请求
type UserQueryReq struct {
	LikeName string `json:"like_name"`
	Age      int    `json:"age"`
}

// Response 统一响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// find 模拟查询
func find(w http.ResponseWriter, r *http.Request) {
	// 只允许 POST 请求
	if r.Method != "POST" {
		// 设置请求头
		w.Header().Set("Content-Type", "application/json")
		resp := Response{
			Code:    http.StatusBadRequest,
			Message: "只允许 POST 请求！",
			Data:    nil,
		}
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			log.Printf("响应写入失败: %v\n", err)
			return
		}
	}
	// 处理 application/json 格式的POST请求携带的数据
	var req UserQueryReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("JSON解析失败: %v\n", err)
		return
	}
	log.Printf("收到查询请求: LikeName包含=%s, Age=%d", req.LikeName, req.Age)
	// 创建一个空切片用来存放结果
	var result []User
	for _, user := range users {
		// 逻辑：如果用户的名字包含 "张" (Strings.Contains)
		// 这里的逻辑可以根据你的需求调整，比如是否同时需要满足 Age 相等
		if strings.Contains(user.Name, req.LikeName) {
			result = append(result, user)
		}
	}
	// 响应结果
	resp := Response{
		Code:    http.StatusOK,
		Message: "查询成功！",
		Data:    result,
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Printf("响应写入失败: %v\n", err)
		return
	}
}

func main() {
	http.Handle("/api/v1/user", http.HandlerFunc(find))
	log.Println("服务端启动成功！")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("服务端启动失败: %v\n", err)
	}
}
