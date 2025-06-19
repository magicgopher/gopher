package abstractfactory

// Cake 实现Product接口的蛋糕食品
type Cake struct{}

// GetName 返回蛋糕名称
func (c *Cake) GetName() string {
	return "Cake"
}

// GetPrice 返回蛋糕价格
func (c *Cake) GetPrice() float64 {
	return 5.0
}

// Biscuit 实现Product接口的饼干食品
type Biscuit struct{}

// GetName 返回饼干名称
func (b *Biscuit) GetName() string {
	return "Biscuit"
}

// GetPrice 返回饼干价格
func (b *Biscuit) GetPrice() float64 {
	return 2.5
}
