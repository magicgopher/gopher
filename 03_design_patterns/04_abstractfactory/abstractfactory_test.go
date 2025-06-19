package abstractfactory

import (
	"testing"
)

// TestCafeFactory 测试咖啡店工厂
func TestCafeFactory(t *testing.T) {
	factory := &CafeFactory{}

	// 测试创建饮品
	drink := factory.CreateDrink()
	coffee, ok := drink.(*Coffee)
	if !ok {
		t.Errorf("期望创建Coffee类型，但得到 %T", drink)
	}
	if name := coffee.GetName(); name != "Coffee" {
		t.Errorf("期望饮品名称为 Coffee，但得到 %s", name)
	}
	if price := coffee.GetPrice(); price != 4.0 {
		t.Errorf("期望饮品价格为 4.0，但得到 %f", price)
	}

	// 测试创建食品
	food := factory.CreateFood()
	cake, ok := food.(*Cake)
	if !ok {
		t.Errorf("期望创建Cake类型，但得到 %T", food)
	}
	if name := cake.GetName(); name != "Cake" {
		t.Errorf("期望食品名称为 Cake，但得到 %s", name)
	}
	if price := cake.GetPrice(); price != 5.0 {
		t.Errorf("期望食品价格为 5.0，但得到 %f", price)
	}
}

// TestTeaShopFactory 测试茶肆工厂
func TestTeaShopFactory(t *testing.T) {
	factory := &TeaShopFactory{}

	// 测试创建饮品
	drink := factory.CreateDrink()
	tea, ok := drink.(*Tea)
	if !ok {
		t.Errorf("期望创建Tea类型，但得到 %T", drink)
	}
	if name := tea.GetName(); name != "Tea" {
		t.Errorf("期望饮品名称为 Tea，但得到 %s", name)
	}
	if price := tea.GetPrice(); price != 3.0 {
		t.Errorf("期望饮品价格为 3.0，但得到 %f", price)
	}

	// 测试创建食品
	food := factory.CreateFood()
	biscuit, ok := food.(*Biscuit)
	if !ok {
		t.Errorf("期望创建Biscuit类型，但得到 %T", food)
	}
	if name := biscuit.GetName(); name != "Biscuit" {
		t.Errorf("期望食品名称为 Biscuit，但得到 %s", name)
	}
	if price := biscuit.GetPrice(); price != 2.5 {
		t.Errorf("期望食品价格为 2.5，但得到 %f", price)
	}
}
