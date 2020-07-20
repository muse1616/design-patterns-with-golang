package main

func main() {
	shape := NewShape("Square")
	//画正方形
	shape.draw()
	shape = NewShape("Circle")
	//画圆形
	shape.draw()
}
