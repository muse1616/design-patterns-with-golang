package main

//装饰器模式 本例采用咖啡 可以添加自定义配料 计算咖啡总价格

// 基类
type Beverage interface {
	//	计算价格
	cost() float64
	getDescription() string
}

//浓缩咖啡
type Espresso struct {
}

func NewEspresso() *Espresso {
	return &Espresso{}
}

//实现基类方法
func (e *Espresso) cost() float64 {
	return 1.99
}
func (e *Espresso) getDescription() string {
	return "Espresso"
}

//综合咖啡
type HouseBlend struct {
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{}
}

//实现基类的方法
func (h *HouseBlend) cost() float64 {
	return 0.89
}
func (h *HouseBlend) getDescription() string {
	return "HouseBlend"
}

//调料
type Mocha struct {
	beverage Beverage
}

//构造
func NewMocha(beverage Beverage) *Mocha {
	return &Mocha{beverage: beverage}
}

//在自身价格基础上增加
func (mocha *Mocha) cost() float64 {
	return 0.20 + mocha.beverage.cost()
}
func (mocha *Mocha) getDescription() string {
	return mocha.beverage.getDescription() + ",Mocha"
}

//糖
type Sugar struct {
	beverage Beverage
}

func NewSugar(beverage Beverage) *Sugar {
	return &Sugar{beverage: beverage}
}
func (sugar *Sugar) cost() float64 {
	return 0.1 + sugar.beverage.cost()
}
func (sugar *Sugar) getDescription() string {
	return sugar.beverage.getDescription() + ",Sugar"
}
