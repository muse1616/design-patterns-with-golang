package main

import "fmt"

// 简单工厂模式
type Shape interface {
	draw()
}
type Circle struct {
}
type Square struct {
}

func (c *Circle) draw() {
	fmt.Println("画圆形")
}
func (s *Square) draw() {
	fmt.Println("画正方形")
}

func NewShape(shapeType string) Shape {
	if shapeType == "Circle" {
		return &Circle{}
	} else if shapeType == "Square" {
		return &Square{}
	}
	return nil
}
