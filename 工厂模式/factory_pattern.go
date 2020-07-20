package main

import "fmt"

//形状
type Shape interface {
	draw()
}

//实现类
type Square struct {
}

func (s Square) draw() {
	fmt.Println("画个正方形")
}

type Circle struct {
}

func (c Circle) draw() {
	fmt.Println("画个圆形")
}

//抽象工厂
type IFactory interface {
	getAShape() Shape
}

//
type SquareFactory struct {
}

func (sf SquareFactory) getAShape() Shape {
	return &Square{}
}

type CircleFactory struct {
}

func (cf CircleFactory) getAShape() Shape {
	return &Circle{}
}
