package main

import "fmt"

func main() {
	/*
		指针（pointer）：指针是一个特殊的变量类型，它的值是某个变量的内存地址。通过指针，你可以间接访问或修改该地址对应的变量值，而不是直接操作变量本身。

		声明指针：使用 * 符号。例如：*int表示一个指向整数的指针。

		获取变量的地址：使用 & 运算符。例如：&x返回变量x的内存地址。

		解引用：使用 * 运算符访问或修改指针指向的值。例如，*p 表示指针 p 所指向的值。
	*/

	x := 10
	var p *int = &x            // p 是一个指向 x 的指针
	fmt.Println("x 的值:", x)    // 输出: 10
	fmt.Println("p 指向的地址:", p) // 输出: x 的内存地址
	fmt.Println("p 指向的值:", *p) // 输出: 10

	*p = 20                     // 通过指针修改 x 的值
	fmt.Println("修改后 x 的值:", x) // 输出: 20
}
