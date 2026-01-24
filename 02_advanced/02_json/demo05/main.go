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
	Address []string `json:"address"` // 字段小写，不会被序列化
}

func main() {
	// 将定义好的切片【切片存储着多个User结构体数据】序列化为json字节切片
	//f1()
	// 将json字符串（字节切片）反序列化为结构体
	f2()
}

func f1() {
	// 定义切片
	users := make([]User, 0, 100)
	// 定义User结构体
	u1 := User{ID: 1, Name: "小刘", Age: 17, Sex: "男", Address: []string{"广州", "珠海", "深圳"}}
	u2 := User{ID: 2, Name: "小红", Age: 18, Sex: "女", Address: []string{"成都", "绵阳", "自贡"}}
	u3 := User{ID: 3, Name: "小米", Age: 19, Sex: "男", Address: []string{"株洲", "湘潭", "永州"}}
	u4 := User{ID: 3, Name: "小美", Age: 20, Sex: "女", Address: []string{"芜湖", "蚌埠", "淮南"}}
	// 将User结构体添加到切片中
	users = append(users, u1, u2, u3, u4)
	// 将结构体反序列化为字节切片
	bytes, err := json.Marshal(&users)
	if err != nil {
		log.Fatalf("JSON 序列化失败: %v", err)
	}
	// 美化输出（带缩进的格式，更易读）
	var prettyJSON []byte
	prettyJSON, err = json.MarshalIndent(&users, "", "\t")
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
	jsonStr := `[{"id":1,"name":"小刘","age":17,"sex":"男","address":["广州","珠海","深圳"]},{"id":2,"name":"小红","age":18,"sex":"女","address":["成都","绵阳","自贡"]},{"id":3,"name":"小米","age":19,"sex":"男","address":["株洲","湘潭","永州"]},{"id":3,"name":"小美","age":20,"sex":"女","address":["芜湖","蚌埠","淮南"]}]`
	// 转换成字节切片
	bytes := []byte(jsonStr)
	// 将字节切片反序列化为User结构体
	users := make([]User, 0, 100)
	err := json.Unmarshal(bytes, &users)
	if err != nil {
		log.Fatalf("JSON 反序列化失败: %v", err)
	}
	for _, user := range users {
		fmt.Printf("反序列化成功，ID=%v, 姓名=%v, 年龄=%v, 性别=%v, 地址=%v\n", user.ID, user.Name, user.Age, user.Sex, user.Address)
	}
}
