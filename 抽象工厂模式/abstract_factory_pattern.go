package main

import "fmt"

//产品族
type BenzCar interface {
	getType()
}

//具体产品
type BenzSportCar struct {
}

func (bsc BenzSportCar) getType() {
	fmt.Println("运动系列奔驰")
}

type BenzBusinessCar struct {
}

func (bc BenzBusinessCar) getType() {
	fmt.Println("商务系列奔驰")
}

type BMWCar interface {
	getType()
}

//具体产品
type BMWSportCar struct {
}

func (bmwSC BMWSportCar) getType() {
	fmt.Println("运动系列宝马")
}

type BMWBusinessCar struct {
}

func (bmwBC BMWBusinessCar) getType() {
	fmt.Println("商务系列宝马")
}

//抽象工厂
type IFactory interface {
	createBenzCar(carType string) BenzCar
	createBMWCar(catType string) BMWCar
}

//具体工厂
type SportFactory struct {
}

func (sf SportFactory) createBenzCar(carType string) BenzCar {
	return &BenzSportCar{}
}
func (sf SportFactory) createBMWCar(carType string) BMWCar {
	return &BMWSportCar{}
}

type BusinessFactory struct {
}

func (bf BusinessFactory) createBenzCar(carType string) BenzCar {
	return &BenzBusinessCar{}
}
func (bf BusinessFactory) createBMWCar(carType string) BMWCar {
	return &BMWBusinessCar{}
}
