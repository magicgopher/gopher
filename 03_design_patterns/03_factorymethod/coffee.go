package factorymethod

// 工厂方法生产的产品

// Coffee 定义咖啡接口
type Coffee interface {
	// GetName 获取咖啡名称
	GetName() string
	// GetPrice 获取咖啡价格
	GetPrice() float64
}

// Latte 实现咖啡接口的拿铁咖啡
type Latte struct{}

// GetName 返回拿铁咖啡的名称
func (l *Latte) GetName() string {
	return "Latte"
}

// GetPrice 返回拿铁咖啡的价格
func (l *Latte) GetPrice() float64 {
	return 4.5
}

// Americano 实现咖啡接口的美式咖啡
type Americano struct{}

// GetName 返回美式咖啡的名称
func (a *Americano) GetName() string {
	return "Americano"
}

// GetPrice 返回美式咖啡的价格
func (a *Americano) GetPrice() float64 {
	return 3.5
}
