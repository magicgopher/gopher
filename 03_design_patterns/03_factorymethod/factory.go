package factorymethod

// 工厂方法的生产产品的具体工厂

// CoffeeFactory 定义咖啡工厂接口
type CoffeeFactory interface {
	// CreateCoffee 创建咖啡实例
	CreateCoffee() Coffee
}

// LatteFactory 实现咖啡工厂接口的拿铁咖啡工厂
type LatteFactory struct{}

// CreateCoffee 创建拿铁咖啡实例
func (lf *LatteFactory) CreateCoffee() Coffee {
	return &Latte{}
}

// AmericanoFactory 实现咖啡工厂接口的美式咖啡工厂
type AmericanoFactory struct{}

// CreateCoffee 创建美式咖啡实例
func (af *AmericanoFactory) CreateCoffee() Coffee {
	return &Americano{}
}
