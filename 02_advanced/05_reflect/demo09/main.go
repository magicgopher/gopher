package main

import (
	"fmt"
	"reflect"
)

type MyInt int

type User struct {
	IdOrName any   `json:"idOrName"`
	M        MyInt `json:"m"`
}

func main() {
	str := 1
	u := User{IdOrName: str, M: 20}
	v := reflect.ValueOf(&u).Elem()

	fmt.Println(u)
	fmt.Println("===============")

	elem1 := v.FieldByName("IdOrName")
	elem2 := v.FieldByName("M")

	fmt.Println(elem1.Type())
	fmt.Println(elem1.Kind())
	fmt.Println(elem2.Type())
	fmt.Println(elem2.Kind())
	fmt.Println("===============")

	if elem1.Kind() == reflect.Interface {
		if elem1.Elem().Kind() == reflect.String {
			elem1.Set(reflect.ValueOf("这是string类型"))
		}
		if elem1.Elem().Kind() == reflect.Int ||
			elem1.Elem().Kind() == reflect.Int32 ||
			elem1.Elem().Kind() == reflect.Int64 {
			elem1.Set(reflect.ValueOf(1000))
		}
	}

	fmt.Println(u)
	fmt.Println("===============")

	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println("name", field.Name)
		fmt.Println("type", field.Type)
		fmt.Println("kind", field.Type.Kind())
		fmt.Println("tag", field.Tag)
	}
}
