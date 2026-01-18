package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 将json字符串转为byte字节切片
	bytes := []byte(`{"age":22,"name":"MagicGopher","sex":"男"}`)
	// 创建一个空的 map，用来存放解析后的 JSON 数据
	m := make(map[string]interface{})
	// 把 JSON 字节数据解析（反序列化）到 map 变量 m 中
	if err := json.Unmarshal(bytes, &m); err != nil {
		log.Fatalf("JSON 解析失败: %v\n", err)
	}
	// 解析成功后 map 里就有数据了，可以正常使用
	fmt.Printf("解析成功！姓名: %v, 年龄: %v, 性别: %v\n", m["name"], m["age"], m["sex"])
}
