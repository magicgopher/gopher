package main

import "fmt"

var a = 100 // 全局变量
var b float64 = 6.66

//c := 200 // 不能使用简短定义的方式定义全局变量

func main() {
	// 变量的注意事项
	// 1.变量必须定义才能使用
	// 2.变量名称不能冲突（同一个作用域内变量名不能冲突）
	// 3.简短定义方式，左边的变量名至少有一个是新的
	// 4.变量的零值，也叫默认值（是数据类型的的默认值）
	// 5.变量定义了就要使用，否则无法通过编译
	var num int
	num = 100
	fmt.Printf("num的值是: %v, 地址是: %p\n", num, &num)
	num = 200
	fmt.Printf("num的值是: %v, 地址是: %p\n", num, &num)
	//fmt.Println(num2) //undefined: num2
	var name string // name定义了字符串类型
	//name = 9.99     // 但是赋值是浮点类型
	//fmt.Println(name) // cannot use 9.99 (untyped float constant) as string value in assignment
	name = "张三"
	fmt.Println(name)
	//num, name := 300, "张三丰" // no new variables on left side of :=
	// := 这个表示简短定义，表示定义变量和赋值变量
	// 使用简短定义变量时，左边的变量必须是新的变量才行
	num, name, money := 300, "张三丰", 9.99
	fmt.Println(num, name, money)
	fmt.Println("=============")
	var m int // 默认值: 0
	fmt.Println(m)
	var n float64 // 默认值: 0.0 -> 0
	fmt.Println(n)
	var s string // 默认值: ""
	fmt.Println(s)
	var s1 []int // 默认值: []
	fmt.Println(s1)
	fmt.Println(s1 == nil)
}
