package main

import "fmt"

// 使用duck作为例子 本例创建一个可以改变鸭子飞行行为的类

//算法接口
type FlyBehavior interface {
	fly() string
}

//算法接口实现类
type FlyWithWings struct {
}
type FlyWithNoWay struct {
}

func (fw *FlyWithWings) fly() string {
	return "使用翅膀飞"
}
func (fnw *FlyWithNoWay) fly() string {
	return "无法飞行"
}

//具体实现类 策略类
type Duck struct {
	name        string
	flyBehavior FlyBehavior
}

//构造函数
func NewDuck(name string, flyBehavior FlyBehavior) *Duck {
	return &Duck{name: name, flyBehavior: flyBehavior}
}

//设置算法类
func (d *Duck) setFlyBehavior(flyBehavior FlyBehavior) {
	d.flyBehavior = flyBehavior
}

//策略类操作方法
func (d *Duck) performFly() {
	fmt.Println(d.name, d.flyBehavior.fly())
}
