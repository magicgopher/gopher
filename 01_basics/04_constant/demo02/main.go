package main

import "fmt"

func main() {
	// 常量的注意事项
	// 1.常量定义后不使用，并不会和变量一样会编译不通过
	// 2.常量赋值之后，是不可以改变的
	// 3.常量中的数据类型只可以是布尔类型、数值型（整数类型、浮点类型、复数类型）和字符类型、字符串类型
	const URL = "https://www.baidu.com"
	fmt.Println(URL)
	//URL = "https://google.com" // annot assign to URL (neither addressable nor a map index expression)
	//fmt.Println(URL)
	const c1 = true // 布尔类型常量
	fmt.Println(c1)
	const c2 = "Hello World!" // 字符串类型常量
	fmt.Println(c2)
	const c3 = 200 // 数值类型（整数）常量
	fmt.Println(c3)
	const c4 = 9.99 // 数值类型（浮点数）常量
	fmt.Println(c4)
	const c5 = 'D' // 字符 rune 常量
	fmt.Println(c5)
	const c6 = '你' // 字符 rune 常量
	fmt.Println(c6)
	//const c7 = [5]int{} // 数组不能作为常量值
	//fmt.Println(c7)
	// 常量定义没有使用，也是可以编译通过的
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
