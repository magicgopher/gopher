package main

import (
	"fmt"
	"reflect"
)

// 以下是一个综合示例，展示如何使用反射分析结构体、修改字段值并调用方法。

// 结构体
type Person struct {
	Name string `tag:"name"`
	Age  int    `tag:"age"`
}

// 方法
func (p Person) Greet(greeting string) string {
	return fmt.Sprintf("%s, %s! You are %d years old.", greeting, p.Name, p.Age)
}

// 函数
func inspectStruct(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	// 处理指针
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("Not a struct")
		return
	}

	// 打印字段信息
	fmt.Println("Fields:")
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		tag := field.Tag.Get("tag")
		fmt.Printf("Field: %s, Value: %v, Tag: %s\n", field.Name, value.Interface(), tag)
	}

	// 修改字段（需可寻址）
	if val.CanSet() {
		ageField := val.FieldByName("Age")
		if ageField.IsValid() && ageField.CanSet() {
			ageField.SetInt(35)
		}
	}

	// 动态调用方法
	method := val.MethodByName("Greet")
	if method.IsValid() {
		args := []reflect.Value{reflect.ValueOf("Hello")}
		results := method.Call(args)
		fmt.Println("Method Result:", results[0].String())
	}
}

func main() {
	p := &Person{Name: "Alice", Age: 30}
	inspectStruct(p)
	fmt.Println("After modification:", p)
}
