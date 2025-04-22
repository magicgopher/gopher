package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := Person{Name: "MagicGopher"}
	v := reflect.ValueOf(p)
	method := v.MethodByName("SayHello")
	args := []reflect.Value{reflect.ValueOf("Hi")}
	result := method.Call(args)
	fmt.Println(result[0].String()) // Hi, MagicGopher
}

// Person 结构体
type Person struct {
	Name string
}

func (p Person) SayHello(greeting string) string {
	return greeting + ", " + p.Name
}
