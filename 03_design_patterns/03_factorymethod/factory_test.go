package factorymethod

import "testing"

// TestLatteFactory 测试拿铁咖啡工厂
func TestLatteFactory(t *testing.T) {
	// 创建拿铁咖啡工厂
	factory := &LatteFactory{}
	// 使用工厂创建咖啡
	coffee := factory.CreateCoffee()

	// 检查咖啡是否为Latte类型
	latte, ok := coffee.(*Latte)
	if !ok {
		t.Errorf("期望创建Latte类型，但得到 %T", coffee)
	}

	// 检查咖啡名称
	if name := latte.GetName(); name != "Latte" {
		t.Errorf("期望名称为 Latte，但得到 %s", name)
	}

	// 检查咖啡价格
	if price := latte.GetPrice(); price != 4.5 {
		t.Errorf("期望价格为 4.5，但得到 %f", price)
	}
}

// TestAmericanoFactory 测试美式咖啡工厂
func TestAmericanoFactory(t *testing.T) {
	// 创建美式咖啡工厂
	factory := &AmericanoFactory{}
	// 使用工厂创建咖啡
	coffee := factory.CreateCoffee()

	// 检查咖啡是否为Americano类型
	americano, ok := coffee.(*Americano)
	if !ok {
		t.Errorf("期望创建Americano类型，但得到 %T", coffee)
	}

	// 检查咖啡名称
	if name := americano.GetName(); name != "Americano" {
		t.Errorf("期望名称为 Americano，但得到 %s", name)
	}

	// 检查咖啡价格
	if price := americano.GetPrice(); price != 3.5 {
		t.Errorf("期望价格为 3.5，但得到 %f", price)
	}
}
