package abstractfactory

// Coffee 实现Product接口的咖啡饮品
type Coffee struct{}

// GetName 返回咖啡名称
func (c *Coffee) GetName() string {
	return "Coffee"
}

// GetPrice 返回咖啡价格
func (c *Coffee) GetPrice() float64 {
	return 4.0
}

// Tea 实现Product接口的茶饮品
type Tea struct{}

// GetName 返回茶名称
func (t *Tea) GetName() string {
	return "Tea"
}

// GetPrice 返回茶价格
func (t *Tea) GetPrice() float64 {
	return 3.0
}
