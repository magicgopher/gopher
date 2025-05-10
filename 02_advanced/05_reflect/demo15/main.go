package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  uint
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello, ", msg)
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 性别: %s\n", p.Name, p.Age, p.Sex)
}

func main() {
	p1 := Person{"Gopher", 18, "未知"}

	DoFileAndMethod(p1)
}

func DoFileAndMethod(input interface{}) {

	// 先获取 input 的类型
	getType := reflect.TypeOf(input)
	fmt.Println("get type is:", getType.Name()) // Person
	fmt.Println("get kind is:", getType.Kind()) // struct

	// 再获取 input 的值
	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue) // {Gopher 18 未知}

	// 获取字段
	// 1. 先获取 interface 的 reflect.Type，然后通过 NumField 进行遍历
	// 2. 再通过 reflect.Type 的 Field 获取其 Field
	// 3. 最后通过 Field 的 Interface() 得到对应的 value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface() //获取第 i 个值
		fmt.Printf("字段名称：%s，字段类型：%s，字段数值：%v\n", field.Name, field.Type, value)
	}

	// 操作方法
	// 1. 先获取 interface 的 reflect.Type，然后通过 .NumMethod 进行遍历
	// 2. 再通过 reflect.Type 的 Method 获取其 Method
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称：%s，方法类型：%v\n", method.Name, method.Type)
	}
}
