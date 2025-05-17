package main

import "fmt"

// 目前Go的方法不支持泛型

//type User struct {
//}

// 编译错误，方法不支持泛型
//func (u User) Hello[T int | string](s T) {
//	fmt.Println(s)
//}

type User[T int | string] struct {
}

func (u *User[T]) Hello(s T) {
	fmt.Println(s)
}

func main() {
	var u1 User[int]
	u1.Hello(100)

	var u2 User[string]
	u2.Hello("hello")
}
