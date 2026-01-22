package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// 客户端示例
// 向服务端发起POST请求，携带表单数据（application/x-www-form-urlencoded）

func main() {
	// 构建POST请求的表单form
	form := url.Values{}
	form.Set("username", "examples@gmail.com")
	form.Set("password", "123456")
	url := "http://127.0.0.1:8080/api/v1/login"
	// 发送POST请求
	resp, err := http.PostForm(url, form)
	if err != nil {
		log.Printf("")
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("响应体关闭失败: %v", err)
		} else {
			log.Println("响应体关闭成功！")
		}
	}()
	// 打印基本信息
	fmt.Printf("HTTP 状态码: %s\n", resp.Status)
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// 读取并打印响应内容（最重要的一步）
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}
	fmt.Println("\n响应内容:")
	fmt.Println(string(body))
	// 判断是否登录成功（简单示例）
	if resp.StatusCode == http.StatusOK {
		fmt.Println("\n请求成功（HTTP 200）")
	} else {
		fmt.Printf("\n请求失败，状态码: %d\n", resp.StatusCode)
	}
}
