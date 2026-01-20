package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// 客户端

func main() {
	// 发起GET请求
	// 查询参数
	// http://127.0.0.1:8080/api/v1/user?name=xxx=age=xx=address=xx
	//queryParameters1()
	//queryParameters2()
	// 路径参数
	// http://127.0.0.1:8080/api/v1/user/{name}/{age}/{address}
	pathParameters()
}

// queryParameters1 使用字符串拼接方式构建GET请求查询参数
func queryParameters1() {
	// 定义GET请求的查询参数
	name := "小李"
	age := 22
	address := "上海"
	// 构建请求的URL
	fullURL := fmt.Sprintf("http://127.0.0.1:8080/api/v1/user?name=%s&age=%d&address=%s", name, age, address)
	fmt.Printf("请求 URL: %s\n", fullURL)
	// 发起GET请求
	resp, err := http.Get(fullURL)
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
	// 检查http状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("服务器返回非 200 状态码: %d\nBody: %s", resp.StatusCode, string(body))
		return
	}
	// 读取响应 body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}
	// 定义响应结构体
	response := struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{}
	// 解析 JSON
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v\n原始内容: %s", err, string(body))
	}
	// 打印结果
	fmt.Println("\n┌──────────────────────────────┐")
	fmt.Printf("  响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("  服务器消息: %s\n", response.Msg)
	fmt.Printf("  返回代码  : %d\n", response.Code)
	fmt.Println("└──────────────────────────────┘")
	// 打印 data 内容（这里 data 是字符串）
	if msg, ok := response.Data.(string); ok {
		fmt.Println("\n返回内容:")
		fmt.Println("  " + msg)
	} else {
		// 如果服务端返回的是结构体，也可以在这里处理
		fmt.Printf("\nData (非字符串类型): %+v\n", response.Data)
	}
	// 额外：如果想直接查看原始 JSON
	fmt.Println("\n原始 JSON 响应:")
	fmt.Println(string(body))
}

// queryParameters2 使用url.Values{}构建GET请求查询参数
func queryParameters2() {
	// 使用url.Values{}来构建查询参数
	params := url.Values{}
	params.Add("name", "小李")
	params.Add("age", "22")
	params.Add("address", "上海")
	// 构建查询的URL
	baseURL := "http://127.0.0.1:8080/api/v1/user"
	fullURL := baseURL + "?" + params.Encode()
	fmt.Printf("请求 URL: %s\n", fullURL)
	// 发起GET请求
	resp, err := http.Get(fullURL)
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
	// 检查http状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("服务器返回非 200 状态码: %d\nBody: %s", resp.StatusCode, string(body))
		return
	}
	// 读取响应 body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}
	// 定义响应结构体
	response := struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{}
	// 解析 JSON
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v\n原始内容: %s", err, string(body))
	}
	// 打印结果
	fmt.Println("\n┌──────────────────────────────┐")
	fmt.Printf("  响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("  服务器消息: %s\n", response.Msg)
	fmt.Printf("  返回代码  : %d\n", response.Code)
	fmt.Println("└──────────────────────────────┘")
	// 打印 data 内容（这里 data 是字符串）
	if msg, ok := response.Data.(string); ok {
		fmt.Println("\n返回内容:")
		fmt.Println("  " + msg)
	} else {
		// 如果服务端返回的是结构体，也可以在这里处理
		fmt.Printf("\nData (非字符串类型): %+v\n", response.Data)
	}
	// 额外：如果想直接查看原始 JSON
	fmt.Println("\n原始 JSON 响应:")
	fmt.Println(string(body))
}

// pathParameters 构建路径参数请求
// 目标 URL: http://127.0.0.1:8080/api/v1/user/{name}/{age}/{address}
func pathParameters() {
	// 定义参数
	name := "小李"
	age := 22
	address := "上海"
	// 构建 URL
	// 【重要】：路径中的参数必须使用 url.PathEscape 进行转义
	// 注意：url.Values{}.Encode() 是用于查询参数的（空格转为+），而 url.PathEscape 是用于路径的（空格转为%20）
	encodedName := url.PathEscape(name)
	encodedAddress := url.PathEscape(address)
	// 使用 Sprintf 按照服务端定义的路由规则拼接
	// 假设服务端路由为 /api/v1/user/{name}/{age}/{address}
	baseURL := "http://127.0.0.1:8080/api/v1/user"
	fullURL := fmt.Sprintf("%s/%s/%d/%s", baseURL, encodedName, age, encodedAddress)
	fmt.Printf("请求 URL: %s\n", fullURL)
	// 发起GET请求
	resp, err := http.Get(fullURL)
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
	// 检查http状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("服务器返回非 200 状态码: %d\nBody: %s", resp.StatusCode, string(body))
		return
	}
	// 读取响应 body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}
	// 定义响应结构体
	response := struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{}
	// 解析 JSON
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v\n原始内容: %s", err, string(body))
	}
	// 打印结果
	fmt.Println("\n┌──────────────────────────────┐")
	fmt.Printf("  响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("  服务器消息: %s\n", response.Msg)
	fmt.Printf("  返回代码  : %d\n", response.Code)
	fmt.Println("└──────────────────────────────┘")
	// 打印 data 内容
	if msg, ok := response.Data.(string); ok {
		fmt.Println("\n返回内容:")
		fmt.Println("  " + msg)
	} else {
		fmt.Printf("\nData (非字符串类型): %+v\n", response.Data)
	}
	// 额外：查看原始 JSON
	fmt.Println("\n原始 JSON 响应:")
	fmt.Println(string(body))
}
