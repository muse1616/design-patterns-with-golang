package main

func main() {
	//var redCircle Shape
	redCircle := NewCircle(100, 20, 20, new(RedCircle))
	blueCircle := NewCircle(20, 30, 30, new(BlueCircle))
	redCircle.draw()
	blueCircle.draw()

}

