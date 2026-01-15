package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// 发起GET请求
	// 请求地址：http://localhost:8080/api/v1/hello?name="xx"&age=xx&address="xx"
	//r1()

	// 发起GET请求
	// 请求地址：http://localhost:8080/api/v1/world/{id}/{name}/{age}/{address}
	r2()
}

func r1() {
	// Query Parameters 查询参数
	name := "张三丰"
	age := 30
	address := "广州"

	// 构建请求的URL地址
	url := fmt.Sprintf("http://localhost:8080/api/v1/hello?name=%s&age=%d&address=%s", name, age, address)

	// 发起请求
	resp, err := http.Get(url)
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

func r2() {
	// Query Parameters 查询参数
	id := 100
	name := "张无忌"
	age := 28
	address := "深圳"

	// 构建请求的URL地址
	url := fmt.Sprintf("http://localhost:8080/api/v1/world/%d/%s/%d/%s", id, name, age, address)

	// 发起请求
	resp, err := http.Get(url)
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
		errorBody, _ := io.ReadAll(resp.Body)
		log.Fatalf("服务器报错 - 状态码: %d, 原因: %s", resp.StatusCode, string(errorBody))
	}

	// 读取服务器返回的响应体内容
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}
	// 将响应体转换为字符串并打印输出
	fmt.Printf("服务器响应: %s\n", string(res))
}
