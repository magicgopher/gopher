package main

import "fmt"

func main() {
	// iota是一个特殊的常量，可以被编译器自动修改的常量
	// 每当定义一个const，iota的初始值为0
	// 每当定义一个常量，就会自动累加1
	// 直到下一个const出现，iota才会清零
	const (
		a = iota    // iota = 0
		b = "hello" // iota = 1 b的值为"hello"
		c = iota    // iota=2
		d           // iota=3 d的值为3
		e = 9.99    // iota=4 e的值为9.99
	)
	fmt.Println(a, b, c, d, e)
	// 下一个const
	const (
		f = iota // 新的const，那么iota又重新从0开始
		i        // iota=1
		j = "你好" // iota=2 j的值: "你好"
		k        // iota=3 k没有值，默认值和上一行一致，k的值: "你好"
		l = 9.99 // iota=4 l的值: 9.99
		m = iota // iota=5
	)
	fmt.Println(f, i, j, k, l)
}
