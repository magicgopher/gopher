package main

import "fmt"

func main() {
	// 变量：计算机中的一小块内存，用于存储数据，在程序运行的过程中是可变的。
	// 以下是定义变量的方式

	// 方式一
	// 定义变量，然后赋值（写在一行或者分开写）
	var v1 int = 100 // 写在一行
	fmt.Printf("v1数据类型: %T, v1变量的值: %v\n", v1, v1)
	var v2 int8 // 定义变量
	v2 = 127    // 再给变量赋值
	fmt.Printf("v2数据类型: %T, v2变量的值: %v\n", v2, v2)

	// 方式二
	// 类型推断，根据等号右边的数值的数据类型，来推断左边变量的数据类型
	var v3 = 5.55
	fmt.Printf("v3的数据类型: %T, v3变量的值: %v\n", v3, v3)
	var v4 = "你好!"
	fmt.Printf("v4数据类型: %T, v4变量的值: %v\n", v4, v4)

	// 方式三
	// 简短定义（简短声明）
	v5 := 999
	fmt.Printf("v5数据类型: %T, v5变量的值: %v\n", v5, v5)

	// 多个变量同时定义
	var a, b, c int // 定义多个变量声明同一种数据类型
	a = 100
	b = 200
	c = 300
	fmt.Println(a, b, c)

	// 多个变量同时定义，使用类型推断变量数据哪种数据类型
	var name, age, money = "张三", 20, 9.99
	fmt.Printf("类型: %T, 值: %v\n", name, name)
	fmt.Printf("类型: %T, 值: %v\n", age, age)
	fmt.Printf("类型: %T, 值: %v\n", money, money)
}
