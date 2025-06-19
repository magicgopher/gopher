package prototype

import "fmt"

// Sheep 定义羊结构体
type Sheep struct {
	// Name 羊的名称
	Name string
	// Age 羊的年龄
	Age int
	// Color 羊的颜色
	Color string
}

// Clone 实现原型模式的克隆方法
// 返回一个新的Sheep实例，复制当前实例的属性
func (s *Sheep) Clone() *Sheep {
	return &Sheep{
		Name:  s.Name,
		Age:   s.Age,
		Color: s.Color,
	}
}

// String 实现Stringer接口，用于打印羊的信息
func (s *Sheep) String() string {
	return fmt.Sprintf("Sheep{Name: %s, Age: %d, Color: %s}", s.Name, s.Age, s.Color)
}

// SetName 设置羊的名称
func (s *Sheep) SetName(name string) {
	s.Name = name
}

// GetName 获取羊的名称
func (s *Sheep) GetName() string {
	return s.Name
}

// SetAge 设置羊的年龄
func (s *Sheep) SetAge(age int) {
	s.Age = age
}

// GetAge 获取羊的年龄
func (s *Sheep) GetAge() int {
	return s.Age
}

// SetColor 设置羊的颜色
func (s *Sheep) SetColor(color string) {
	s.Color = color
}

// GetColor 获取羊的颜色
func (s *Sheep) GetColor() string {
	return s.Color
}
