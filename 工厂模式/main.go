package main

func main() {
	var f IFactory
	f = new(SquareFactory)
	f.getAShape().draw()
	f = new(CircleFactory)
	f.getAShape().draw()
}
