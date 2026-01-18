package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 复杂的数据结构序列化示例

func main() {
	// 序列化切片map
	//f1()
	// 反序列化切片map
	f2()
}

func f1() {
	// 定义3个map
	obj1 := map[string]interface{}{"id": 1, "name": "张三", "age": 17, "sex": "男"}
	obj2 := map[string]interface{}{"id": 2, "name": "李四", "age": 18, "sex": "女"}
	obj3 := map[string]interface{}{"id": 3, "name": "王五", "age": 19, "sex": "男"}
	obj4 := map[string]interface{}{"id": 4, "name": "赵六", "age": 20, "sex": "女"}
	// 定义一个切片用于存储map
	objs := make([]map[string]interface{}, 0, 10)
	objs = append(objs, obj1, obj2, obj3, obj4)
	// 将切片转为 JSON 字节切片
	bytes, err := json.Marshal(&objs)
	if err != nil {
		log.Fatalf("JSON 序列化失败: %v", err)
	}
	// 美化输出（带缩进的格式，更易读）
	var prettyJSON []byte
	prettyJSON, err = json.MarshalIndent(objs, "", "  ")
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
	jsonStr := `[{"age":17,"id":1,"name":"张三","sex":"男"},{"age":18,"id":2,"name":"李四","sex":"女"},{"age":19,"id":3,"name":"王五","sex":"男"},{"age":20,"id":4,"name":"赵六","sex":"女"}]`
	// 将json字符串转为字节切片
	bytes := []byte(jsonStr)
	// 定义接收反序列化结果的变量
	var objs []map[string]interface{}
	// 将字节切片反序列化为go的数据类型，这里就是定义的[]map[string]interface{}
	err := json.Unmarshal(bytes, &objs)
	if err != nil {
		log.Fatalf("JSON 反序列化失败: %v", err)
	}
	// 反序列化成功！现在 objs 就是原来的切片了
	fmt.Println("反序列化成功，共", len(objs), "条记录")
	// 遍历
	for i, obj := range objs {
		fmt.Printf("第 %d 条: ID=%v, 姓名=%v, 年龄=%v, 性别=%v\n",
			i+1, obj["id"], obj["name"], obj["age"], obj["sex"])
	}
}
