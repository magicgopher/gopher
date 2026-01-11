package main

import "github.com/magicgopher/gopher/01_basics/02_package/demo01/example"

func main() {
	// 在main包下使用example包下的SayHello函数
	example.SayHello("MagicGopher")
}
