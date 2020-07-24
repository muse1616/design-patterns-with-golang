package main

import "fmt"

func main() {
	o1 := new(Object)
	o1.id = 1
	o2 := o1.clone()
	o2.id = 2
	// o1 != o2
	fmt.Println(o1, o2)
}
