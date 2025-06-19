package abstractfactory

// Product 接口定义产品行为
type Product interface {
	// GetName 获取产品名称
	GetName() string
	// GetPrice 获取产品价格
	GetPrice() float64
}

// AbstractFactory 接口定义抽象工厂
type AbstractFactory interface {
	// CreateDrink 创建饮品
	CreateDrink() Product
	// CreateFood 创建食品
	CreateFood() Product
}

// CafeFactory 实现抽象工厂的咖啡店工厂
type CafeFactory struct{}

// CreateDrink 创建咖啡饮品
func (cf *CafeFactory) CreateDrink() Product {
	return &Coffee{}
}

// CreateFood 创建蛋糕食品
func (cf *CafeFactory) CreateFood() Product {
	return &Cake{}
}

// TeaShopFactory 实现抽象工厂的茶肆工厂
type TeaShopFactory struct{}

// CreateDrink 创建茶饮品
func (tsf *TeaShopFactory) CreateDrink() Product {
	return &Tea{}
}

// CreateFood 创建饼干食品
func (tsf *TeaShopFactory) CreateFood() Product {
	return &Biscuit{}
}
