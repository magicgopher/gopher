package simplefactory

import "testing"

// TestNewCoffee 测试 NewCoffee 工厂函数的功能
// 验证是否能正确创建 Latte 和 Americano 咖啡实例，并检查其名称和价格
// 同时验证对未知咖啡类型的处理是否返回 nil
func TestNewCoffee(t *testing.T) {
	// 定义测试用例结构体切片，包含测试输入和预期输出
	tests := []struct {
		coffeeType string  // 输入的咖啡类型
		wantName   string  // 预期的咖啡名称
		wantPrice  float64 // 预期的咖啡价格
		wantNil    bool    // 是否预期返回 nil
	}{
		{"latte", "Latte", 12.5, false},        // 测试 Latte 咖啡：应返回非 nil，名称为 "Latte"，价格为 12.5
		{"americano", "Americano", 9.9, false}, // 测试 Americano 咖啡：应返回非 nil，名称为 "Americano"，价格为 9.9
		{"unknown", "", 0.0, true},             // 测试未知类型：应返回 nil
	}

	// 遍历测试用例
	for _, tt := range tests {
		// 为每个测试用例创建一个子测试，名称为输入的咖啡类型
		t.Run(tt.coffeeType, func(t *testing.T) {
			// 调用 NewCoffee 函数创建咖啡实例
			coffee := NewCoffee(tt.coffeeType)

			// 如果预期返回 nil
			if tt.wantNil {
				// 检查 NewCoffee 返回值是否为 nil
				if coffee != nil {
					t.Errorf("NewCoffee(%q) = %v, 期望返回 nil", tt.coffeeType, coffee)
				}
				return
			}

			// 如果返回值为 nil，但预期非 nil，则报错
			if coffee == nil {
				t.Errorf("NewCoffee(%q) = nil, 期望返回非 nil", tt.coffeeType)
				return
			}

			// 检查咖啡名称是否符合预期
			if name := coffee.Name(); name != tt.wantName {
				t.Errorf("Name() = %q, 期望 %q", name, tt.wantName)
			}

			// 检查咖啡价格是否符合预期
			if price := coffee.Price(); price != tt.wantPrice {
				t.Errorf("Price() = %v, 期望 %v", price, tt.wantPrice)
			}
		})
	}
}
