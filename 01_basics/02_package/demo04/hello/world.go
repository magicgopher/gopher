package hello

import "fmt"

func Say() string {
	hello := sayHello("World!") // sayHello()函数名小写所以只能在当前包内访问
	return fmt.Sprintf("Golang, %v", hello)
}
