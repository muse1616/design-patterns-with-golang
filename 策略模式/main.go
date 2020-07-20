package main

func main() {

	d1 := NewDuck("mike", new(FlyWithWings))

	d1.performFly()
	//重置方法
	d1.setFlyBehavior(new(FlyWithNoWay))
	d1.performFly()
}
