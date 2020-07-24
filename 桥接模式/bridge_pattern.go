package main

import "fmt"

// 桥接接口
type DrawAPI interface {
	drawCircle(radius, x, y int)
}

//实现类
type RedCircle struct {
}
type BlueCircle struct {
}

func (rc *RedCircle) drawCircle(radius, x, y int) {
	fmt.Printf("drawing a red circle,radius:%d,x:%d,y:%d\n", radius, x, y)
}
func (bc *BlueCircle) drawCircle(radius, x, y int) {
	fmt.Printf("drawing a blue circle,radius:%d,x:%d,y:%d\n", radius, x, y)
}

// 父类
type Shape struct {
	drawAPI DrawAPI
}

type Circle struct {
	x      int
	y      int
	radius int
	// 继承Shape
	Shape
}

func NewCircle(radius, x, y int, drawAPI DrawAPI) *Circle {
	return &Circle{
		x:      x,
		y:      y,
		radius: radius,
		Shape:  Shape{drawAPI: drawAPI},
	}
}
func (c *Circle) draw() {
	c.drawAPI.drawCircle(c.radius, c.x, c.y)
}
