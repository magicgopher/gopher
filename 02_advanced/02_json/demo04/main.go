package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User 用户结构体
type User struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Age     uint     `json:"age"`
	Sex     string   `json:"sex"`
	address []string `json:"address"` // 字段小写，不会被序列化
}

func main() {
	// 结构体序列化为json字符串（字节切片）
	f1()
	// 将json字符串（字节切片）反序列化为结构体
	//f2()
}

func f1() {
	// 定义结构体
	u := User{
		ID:      1,
		Name:    "张三丰",
		Age:     17,
		Sex:     "男",
		address: []string{"广州", "珠海", "深圳"},
	}
	// 将结构体反序列化为字节切片
	bytes, err := json.Marshal(&u)
	if err != nil {
		log.Fatalf("JSON 序列化失败: %v", err)
	}
	// 美化输出（带缩进的格式，更易读）
	var prettyJSON []byte
	prettyJSON, err = json.MarshalIndent(u, "", "\t")
	if err != nil {
		log.Fatalf("美化JSON失败: %v", err)
	}
	// 打印结果
	fmt.Println("JSON 字符串（紧凑格式）：")
	fmt.Println(string(bytes))
	fmt.Println("\n美化后的 JSON（推荐阅读）：")
	fmt.Println(string(prettyJSON))
}

func f2() {
	// 准备 JSON 字符串（这是 f1 中序列化后的结果）
	jsonStr := `{"id":1,"name":"张三丰","age":17,"sex":"男"}`
	// 转换成字节切片
	bytes := []byte(jsonStr)
	// 将字节切片反序列化为User结构体
	user := &User{}
	err := json.Unmarshal(bytes, &user)
	if err != nil {
		log.Fatalf("JSON 反序列化失败: %v", err)
	}
	fmt.Printf("反序列化成功，ID=%v, 姓名=%v, 年龄=%v, 性别=%v\n", user.ID, user.Name, user.Age, user.Sex)
}
