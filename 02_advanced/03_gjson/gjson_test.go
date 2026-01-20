package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
	"testing"
)

// gjson是一个可以从json结构中拿到某个属性值

const jsonStr = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ],
  "city": ["New York", "Chicago", "Washington", "Boston", "Detroit", "Seattle"]
}`

// TestGet1 Get函数测试
func TestGet1(t *testing.T) {
	value := gjson.GetBytes([]byte(jsonStr), "name")
	t.Log(value.String())
}

// TestGet2 Get函数测试
func TestGet2(t *testing.T) {
	bytes := []byte(jsonStr)
	res1 := gjson.GetBytes(bytes, "age")
	t.Log(res1)
}

// TestGet3 Get函数测试
func TestGet3(t *testing.T) {
	// 层级访问使用 .
	v1 := gjson.Get(jsonStr, "name.first")
	t.Log(v1.String())
	// 数组访问用数字索引
	v2 := gjson.Get(jsonStr, "children.1")
	t.Log(v2)
	// 数组长度使用#
	v3 := gjson.Get(jsonStr, "city.#")
	t.Log(v3)
	// 通配符
	// *匹配零个或多个字符
	v4 := gjson.Get(jsonStr, "child*.2")
	// ?表示匹配任意一个字符
	v5 := gjson.Get(jsonStr, "c?ildren.0")
	t.Log(v4)
	t.Log(v5)
}

// TestGet4 Get函数测试
func TestGet4(t *testing.T) {
	v1 := gjson.Get(jsonStr, "friends.#(last==\"Murphy\").first")
	t.Log(v1)
}

// 定义两个json，一个格式错误、一个格式正确
const jsonStr2 = `{ "name": "haha", age: 1}`
const jsonStr3 = `{ "name": "haha", "age": 1}`

// TestValid 验证这个json格式是否正确
func TestValid(t *testing.T) {
	valid1 := gjson.Valid(jsonStr2)
	t.Log(valid1)
	valid2 := gjson.Valid(jsonStr3)
	t.Log(valid2)
}

const jsonStr4 = `{"name":{"first":"Janet","last":null},"age":47}`

func TestGetExists(t *testing.T) {
	// 检查一个存在的路径
	value := gjson.Get(jsonStr4, "name.first")
	if value.Exists() {
		fmt.Println("First name exists:", value.String())
	} else {
		fmt.Println("First name does not exist")
	}

	// 检查一个不存在的路径
	missing := gjson.Get(jsonStr4, "name.middle")
	if !missing.Exists() {
		fmt.Println("Middle name does not exist")
	}
}
