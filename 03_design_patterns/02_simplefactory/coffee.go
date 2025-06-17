package simplefactory

// Coffee 咖啡接口
type Coffee interface {
	Name() string   // 名称
	Price() float64 // 价格
}

// Latte 拿铁咖啡
type Latte struct {
}

func (l *Latte) Name() string {
	return "Latte"
}

func (l *Latte) Price() float64 {
	return 12.5
}

// Americano 美式咖啡
type Americano struct {
}

func (a *Americano) Name() string {
	return "Americano"
}

func (a *Americano) Price() float64 {
	return 9.9
}
