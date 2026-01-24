package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// Response 定义一个结构体用来解析服务端返回的 JSON
type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"` // Data 里面可能是各种数据，用 map 通吃
}

func main() {
	// 1. 准备要上传的目标 URL
	targetUrl := "http://localhost:8080/api/v1/upload"

	// 2. 准备一个本地文件用于测试 (为了方便演示，我先代码创建一个)
	localFileName := "client_test.txt"
	createDummyFile(localFileName)

	// --- 核心步骤：构建 multipart/form-data 数据包 ---

	// 创建一个缓冲区，用来存放请求体的数据
	bodyBuf := &bytes.Buffer{}
	// 创建 multipart 写入器
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 【关键点】创建文件字段
	// 参数1 "file": 必须和服务端 r.FormFile("file") 里的字符串完全一致！
	// 参数2 localFileName: 告诉服务端这个文件叫什么名字
	fileWriter, err := bodyWriter.CreateFormFile("file", localFileName)
	if err != nil {
		panic(err)
	}

	// 打开本地文件
	fh, err := os.Open(localFileName)
	if err != nil {
		panic(err)
	}
	// 将本地文件内容拷贝到请求体中
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		panic(err)
	}
	fh.Close()

	// 结束 multipart 写入 (必须调用，用于写入结尾的 boundary)
	bodyWriter.Close()

	// --- 发送请求 ---

	// 创建 Request 对象
	req, err := http.NewRequest("POST", targetUrl, bodyBuf)
	if err != nil {
		panic(err)
	}

	// 【关键点】设置 Content-Type
	// 必须包含 boundary，例如: multipart/form-data; boundary=...
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// --- 处理响应 ---

	// 读取响应的所有内容
	respBytes, _ := io.ReadAll(resp.Body)

	// 尝试解析 JSON
	var serverResp Response
	if err := json.Unmarshal(respBytes, &serverResp); err != nil {
		// 如果解析失败，说明服务端可能没返回 JSON，或者报错了
		fmt.Println("解析响应失败，原始内容:", string(respBytes))
		return
	}

	// 打印结果
	fmt.Println("--------------------------------")
	fmt.Printf("状态码: %d\n", serverResp.Code)
	fmt.Printf("服务端消息: %s\n", serverResp.Message)
	if serverResp.Data != nil {
		fmt.Printf("保存的文件名: %v\n", serverResp.Data["filename"])
		fmt.Printf("文件大小: %v 字节\n", serverResp.Data["size"])
	}
	fmt.Println("--------------------------------")
}

// 辅助函数：创建一个用来测试的本地文件
func createDummyFile(filename string) {
	content := "这是客户端生成的测试数据。\nGo语言网络编程非常有意思！"
	_ = os.WriteFile(filename, []byte(content), 0666)
	fmt.Printf("本地测试文件 [%s] 已准备好。\n", filename)
}
