package main

import "fmt"

type Shape interface {
	draw()
}
type Point struct {
	x int
	y int
}
type Circle struct {
	id int
	p  Point
}

func NewCircle(id int, p Point) *Circle {
	return &Circle{
		id: id,
		p:  p,
	}
}

func (c Circle) draw() {
	fmt.Printf("id:%d,x:%d,y:%d", c.id, c.p.x, c.p.y)
}

type ShapeFactory interface {
	getShape(id int) Shape
}

type CircleFactory struct {
	circleMap map[int]*Circle
}

func (c *CircleFactory) getShape(id int) Shape {
	if v, ok := c.circleMap[id]; ok {
		//	存在
		return v
	} else {
		//	先创建一个
		fmt.Println("创建一个圆,id:", id)
		circle := NewCircle(id, Point{id + 1, id - 1})
		c.circleMap[id] = circle
		return circle
	}

}

func NewCircleFactory() *CircleFactory {
	return &CircleFactory{circleMap: make(map[int]*Circle)}
}
