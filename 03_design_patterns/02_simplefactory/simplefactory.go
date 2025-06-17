package simplefactory

// NewCoffee 简单工厂方法，根据类型创建咖啡
func NewCoffee(coffeeType string) Coffee {
	switch coffeeType {
	case "americano":
		return &Americano{}
	case "latte":
		return &Latte{}
	default:
		return nil
	}
}
