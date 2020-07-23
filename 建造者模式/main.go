package main

import "fmt"

func main() {
	//	一个builder
	builder := new(ObjectBuilder)
	director := &Director{builder: builder}
	o := new(Object)
	o = director.Create("1", "muse")
	fmt.Printf("%+v", o)
}
