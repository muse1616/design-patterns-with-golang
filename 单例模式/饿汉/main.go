package main

import "fmt"

func main() {
	// 线程安全
	i := GetIns()
	i.id = 3
	j := GetIns()
	// 3
	fmt.Println(j.id)
}
