package main

import (
	"fmt"
	"reflect"
)

/**
观察者模式 本例采用天气数据布告板例子 天气数据改变使两个布告板随着数据变化而改变
*/

//主题接口 观察者订阅主题 主题接口可以添加观察者 移除观察者 通知所有观察者
type Subject interface {
	//	注册观察者
	registerObserver(o Observer)
	//移除观察者
	removeObserver(o Observer)
	//通知所有观察者
	notifyObservers(data WeatherData)
}
//观察者接口 用来更新观察者数据
type Observer interface {
	update(wd WeatherData)
}
//面板显示接口
type Displayment interface {
	display()
}


//实现主题接口
type WeatherData struct {
	// 实现Subject

	Subject
	//	观察者列表
	observers []Observer
	//	数据
	temperature float64
	humidity    float64
	pressure    float64
}

//构造方法
func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make([]Observer, 0),
	}
}
//实现接口的三个方法
func (wd *WeatherData) registerObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}
func (wd *WeatherData) removeObserver(o Observer) {
	for k, ob := range wd.observers {
		if reflect.DeepEqual(ob, o) {
			//去掉o
			wd.observers = append(wd.observers[:k], wd.observers[k+1:]...)
			break
		}
	}
}
func (wd *WeatherData) notifyObservers(data WeatherData) {
	for _, o := range wd.observers {
		o.update(*wd)
	}
}

//触发更新函数
func (wd *WeatherData) updateMeasurement(temperature float64, humidity float64, pressure float64) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.notifyObservers(*wd)
}

//布告板
type CurrentConditionsDisplay struct {
	//继承显示接口
	Displayment
	//继承观察者
	Observer
	temperature float64
	humidity    float64
}

//构造函数 注册为观察者
func NewCurrentConditionDisplay(s Subject) *CurrentConditionsDisplay {
	c := new(CurrentConditionsDisplay)
	s.registerObserver(c)
	return c
}

//显示布告板
func (c *CurrentConditionsDisplay) display() {
	fmt.Println("Current conditions: ", c.temperature, "F degrees and ", c.humidity, "% humidity")
}

//更新数据
func (c *CurrentConditionsDisplay) update(wd WeatherData) {
	c.temperature = wd.temperature
	c.humidity = wd.humidity
	c.display()
}
