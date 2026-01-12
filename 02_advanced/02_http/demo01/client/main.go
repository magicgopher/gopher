package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// 请求写带的参数
	param := "MagicGopher"

	// 没有携带参数
	//url1 := "http://127.0.0.1:8080/hello"
	url2 := fmt.Sprintf("http://127.0.0.1:8080/hello?param=%s", param)

	// 发起 get 请求
	//resp, err := http.Get(url1)
	resp, err := http.Get(url2)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}

	// defer 延迟关闭响应体，并在关闭失败时记录警告
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("关闭响应体失败: %v", err)
		}
	}()

	// 检查 HTTP 状态码是否为 200 OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("服务器返回错误:%v\n", err)
	}

	// 读取服务器返回的响应体内容
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}
	// 将响应体转换为字符串并打印输出
	fmt.Printf("服务器响应: %s\n", string(res))
}
