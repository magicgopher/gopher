package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {
	// 通过反射，来更改对象的值，前提是数值可以被更改
	s1 := Student{"Gopher", 18, "XXX小学"}
	fmt.Printf("s1 类型：%T\n", s1)

	p1 := &s1
	fmt.Printf("p1 类型：%T\n", p1)
	fmt.Println("s1.Name:", s1.Name)
	fmt.Println("*p1.Name:", (*p1).Name)

	// 使用 reflect.ValueOf 取出值，注意传指针
	v1 := reflect.ValueOf(&s1)

	if v1.Kind() == reflect.Ptr {
		fmt.Println("v1.Elem 是否可以设置：", v1.Elem().CanSet())
		v1 = v1.Elem()
	}

	f1 := v1.FieldByName("Name")
	f1.SetString("MagicGopher")
	f2 := v1.FieldByName("Age")
	f2.SetInt(20)
	f3 := v1.FieldByName("School")
	f3.SetString("XXX学校")
	fmt.Println("s1:", s1)
}
