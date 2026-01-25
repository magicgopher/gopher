package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 客户端示例代码
// 发送POST请求携带（application/json）格式数据

// Response 统一响应结构体
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	// 构建响应体的数据
	param := map[string]interface{}{
		"like_name": "张",
		"age":       20,
	}
	// 序列化为json格式数据
	jsonData, err := json.Marshal(&param)
	if err != nil {
		log.Printf("JSON序列化失败: %v\n", err)
		return
	}
	url := "http://127.0.0.1:8080/api/v1/user"
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Printf("请求发送失败: %v\n")
		return
	}
	// defer关闭响应体
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("响应体关闭失败: %v\n", err)
		} else {
			log.Println("响应体关闭成功！")
		}
	}()
	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("服务端返回异常状态码: %d", resp.StatusCode)
	}
	// 读取响应体的内容
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}
	fmt.Printf("服务端原始响应: %s\n", string(bodyBytes))
	var response Response
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		log.Fatalf("解析响应JSON失败: %v", err)
	}
	if dataSlice, ok := response.Data.([]interface{}); ok {
		for _, item := range dataSlice {
			// 2. 再断言切片里的每个元素是 Map
			if itemMap, ok := item.(map[string]interface{}); ok {
				id := int(itemMap["id"].(float64))
				name := itemMap["name"].(string)
				age := int(itemMap["age"].(float64))
				sex := itemMap["sex"].(string)
				fmt.Printf("id: %d, name: %s, age: %d, sex: %s\n", id, name, age, sex)
			}
		}
	}
}
