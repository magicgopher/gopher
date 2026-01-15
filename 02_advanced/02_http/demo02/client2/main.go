package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// 使用 url.Values 来构建query params
	// 例如 http://example.com/hello?name=xxx&age=xxx&address=xxx
	params := url.Values{}
	params.Add("name", "萧火火")
	params.Add("age", "17")
	params.Add("address", "斗气大陆")

	// Encode() 方法会自动处理拼接和特殊字符转义（Url Encode）
	baseURL := "http://localhost:8080/api/v1/hello"
	fullURL := baseURL + "?" + params.Encode()

	// 发起GET请求
	resp, err := http.Get(fullURL)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	// defer 延迟关闭响应体，并在关闭失败时记录警告
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("关闭响应体失败: %v", err)
		}
	}()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("响应: %s\n", body)
}
