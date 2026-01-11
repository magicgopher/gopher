package main

import (
	"fmt"
	"github.com/magicgopher/gopher/01_basics/02_package/demo04/hello"
)

func main() {
	res1 := hello.SayHello("World!") // 访问hello包下的SayHello()函数
	fmt.Println(res1)
	//hello.sayHello("World") // 使用访问hello包下的sayHello()函数

	h := hello.Hello{
		Value: "Hello",
	} // 通过包访问 Hello 结构体
	res2 := h.Hello()
	fmt.Println(res2)
}
