package main

import (
	"fmt"
	"reflect"
)

// 反射示例
// 通用结构体转Map

type Person struct {
	Name string
	Age  uint
}

//func (p Person) SayHello() string {
//	return "Hello, " + p.Name
//}

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" { // 跳过私有字段
			continue
		}
		result[field.Name] = v.Field(i).Interface()
	}
	return result
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	m := StructToMap(p)
	fmt.Println(m)
}
