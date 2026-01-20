package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// 客户端
// 发起http请求（get请求）服务器获取数据

func main() {
	// 定义请求携带的参数
	param := "Hello"
	// 定义请求的url路径
	url := fmt.Sprintf("http://127.0.0.1:8080/hello?value=%s", param)
	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	// defer关闭响应
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("响应关闭失败: %v\n", closeErr)
		} else {
			log.Println("响应关闭成功！")
		}
	}()
	// 检查响应状态码（通常先判断是不是200）
	if resp.StatusCode != http.StatusOK {
		log.Printf("服务器返回了错误状态码: %d %s", resp.StatusCode, resp.Status)
		return
	}
	// 读取服务器返回的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应内容失败: %v", err)
		return
	}
	// 输出响应状态码和服务端响应的内容
	fmt.Printf("服务器返回的状态码: %v\n", resp.StatusCode)
	fmt.Printf("服务器返回的内容: %v\n", string(body))
}
