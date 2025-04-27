package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.14
	// 第一定律：反射从接口值转变为反射对象
	fmt.Println("type:", reflect.TypeOf(x))

	// 第二定律：反射从反射对象转变为接口值
	v := reflect.ValueOf(x)
	y := v.Interface().(float64)
	fmt.Println("y=", y)

	// 第三定律：要修改反射对象的值，其值必须可以设置
	//v.SetFloat(7.1) // panic: reflect: reflect.Value.SetFloat using unaddressable value
	pv := reflect.ValueOf(&x)
	elem := pv.Elem()
	elem.SetFloat(7.1)
	fmt.Println("x=", x)
}
