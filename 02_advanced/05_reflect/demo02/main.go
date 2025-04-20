package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	t := reflect.TypeOf(p)

	fmt.Println("Type Name:", t.Name()) // Type Name: Person
	fmt.Println("Kind:", t.Kind())      // Kind: struct

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: Name=%s, Type=%s\n", i, field.Name, field.Type)
	}
}
