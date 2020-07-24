package main

import "fmt"

type Shape interface {
	draw()
}
type Circle struct {
}

func (c *Circle) draw() {
	fmt.Println("draw a Circle")
}

type Square struct {
}

func (s *Square) draw() {
	fmt.Println("draw a Square")
}

type ShapeMaker struct {
	circle Shape
	square Shape
}

func NewShapeMaker() *ShapeMaker {
	return &ShapeMaker{
		circle: new(Circle),
		square: &Square{},
	}
}
func (sm *ShapeMaker)drawACircle()  {
	sm.circle.draw()
}
func (sm *ShapeMaker)drawASquare()  {
	sm.square.draw()
}
