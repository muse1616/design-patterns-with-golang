package main

import "fmt"

func main() {
	var fc ShapeFactory
	fc = NewCircleFactory()
	fc.getShape(1)
	fc.getShape(1).draw()
	fmt.Println()
	fc.getShape(2)
	// 此时不再新创建
	fc.getShape(2)
	fc.getShape(3)

}
