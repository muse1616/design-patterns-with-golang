package main

func main() {
	var f BusinessFactory
	car := f.createBenzCar("asd")
	car.getType()

	var sf SportFactory
	car = sf.createBMWCar("asd")
	car.getType()
}
