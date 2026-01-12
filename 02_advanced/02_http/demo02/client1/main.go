package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:8080/api/v1/user/"

	req := map[string]interface{}{
		"name": "李白",
		"age":  18,
	}

	// 将map序列化为json字节
	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Fatal("JSON 序列化失败:", err)
	}

	fmt.Printf("发送的 JSON 内容: %s\n", string(jsonData))

	// 发起Post请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("请求发送失败:", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("响应体关闭失败: %v\n", closeErr)
		}
	}()

	// 读取响应体的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("读取响应失败:", err)
	}
	fmt.Println("状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))
}
