package prototype

import (
	"testing"
)

// TestSheepClone 测试羊的克隆功能
func TestSheepClone(t *testing.T) {
	// 创建原始羊
	original := &Sheep{
		Name:  "Dolly",
		Age:   3,
		Color: "White",
	}

	// 克隆羊
	clone := original.Clone()

	// 检查克隆羊是否具有相同的属性
	if clone.GetName() != "Dolly" {
		t.Errorf("期望克隆羊名称为 Dolly，但得到 %s", clone.GetName())
	}
	if clone.GetAge() != 3 {
		t.Errorf("期望克隆羊年龄为 3，但得到 %d", clone.GetAge())
	}
	if clone.GetColor() != "White" {
		t.Errorf("期望克隆羊颜色为 White，但得到 %s", clone.GetColor())
	}

	// 修改克隆羊的属性，验证不影响原始羊
	clone.SetName("Molly")
	clone.SetAge(4)
	clone.SetColor("Black")

	// 检查原始羊的属性未被修改
	if original.GetName() != "Dolly" {
		t.Errorf("原始羊名称应为 Dolly，但得到 %s", original.GetName())
	}
	if original.GetAge() != 3 {
		t.Errorf("原始羊年龄应为 3，但得到 %d", original.GetAge())
	}
	if original.GetColor() != "White" {
		t.Errorf("原始羊颜色应为 White，但得到 %s", original.GetColor())
	}

	// 检查克隆羊的修改是否正确
	if clone.GetName() != "Molly" {
		t.Errorf("克隆羊名称应为 Molly，但得到 %s", clone.GetName())
	}
	if clone.GetAge() != 4 {
		t.Errorf("克隆羊年龄应为 4，但得到 %d", clone.GetAge())
	}
	if clone.GetColor() != "Black" {
		t.Errorf("克隆羊颜色应为 Black，但得到 %s", clone.GetColor())
	}
}
