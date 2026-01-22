package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// 服务端示例
// 处理POST请求（application/x-www-form-urlencoded）格式的数据

type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 信息
	Data    interface{} `json:"data"`    // 数据
}

func Login(w http.ResponseWriter, r *http.Request) {
	// 只接收POST请求
	if r.Method != "POST" {
		// 设置响应头
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp := Response{
			Code:    405,
			Message: "只支持 POST 请求",
			Data:    nil,
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	// 处理POST请求携带的数据
	// application/x-www-form-urlencoded
	if err := r.ParseForm(); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest) // 400
		resp := Response{
			Code:    1001,
			Message: "表单解析失败：" + err.Error(),
			Data:    nil,
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	// 获取表单信息，拿到用户名和密码
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Printf("用户名: %v, 密码: %v\n", username, password)
	// 简单校验：不能为空
	if username == "" || password == "" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK) // 这里很多项目用 200 + 业务错误码
		resp := Response{
			Code:    1002,
			Message: "用户名和密码不能为空",
			Data:    nil,
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	// 模拟登录
	if username == "examples@gmail.com" && password == "123456" {
		// 登录成功
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		resp := Response{
			Code:    0,
			Message: "登录成功",
			Data: map[string]string{
				"token":  "fake-jwt-token-abc123xyz",
				"user":   username,
				"role":   "user",
				"expire": "2026-02-01",
			},
		}
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		resp := Response{
			Code:    1003,
			Message: "用户名或密码错误",
			Data:    nil,
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
}

func main() {
	// 关于http的post请求常见的四种格式
	// application/x-www-form-urlencoded
	// multipart/form-data
	// application/json
	// application/octet-stream
	// 处理 application/x-www-form-urlencoded 格式
	http.Handle("/api/v1/login", http.HandlerFunc(Login))
	log.Println("服务端启动成功！")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("服务端启动失败: %v\n", err)
	}
}
