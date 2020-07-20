package main

import (
	"fmt"
)

func main() {
	//	一杯加糖和摩卡的速溶
	var e1 Beverage
	e1 = NewEspresso()
	e1 = NewMocha(e1)
	e1 = NewSugar(e1)
	//Espresso,Mocha,Sugar  $ 2.29
	fmt.Println(e1.getDescription()," $",e1.cost())
	//双倍摩卡的综合咖啡
	var h1 Beverage
	h1 = NewHouseBlend()
	h1 = NewMocha(h1)
	h1 = NewMocha(h1)
	//HouseBlend,Mocha,Mocha  $ 1.29
	fmt.Println(h1.getDescription()," $",h1.cost())


}
