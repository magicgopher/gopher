package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 序列化: 将内存中的对象转换为一种可以存储或传输的格式（通常是字节流、JSON 字符串或 XML）的过程。
// 反序列化: 将序列化后的数据（字节流、JSON 等）重新恢复为内存中的对象的过程。

func main() {
	// 定义map
	m := map[string]interface{}{"name": "MagicGopher", "age": 22, "sex": "男"}
	// 将 map 序列化为json字节切片
	bytes, err := json.Marshal(&m)
	if err != nil {
		log.Fatalf("json序列化失败: %v\n", err)
	}
	fmt.Printf("json字节切片转为json字符串, 内容: %v\n", string(bytes))
}
