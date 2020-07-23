package main

import (
	"fmt"
	"sync"
)

func main() {
	// 难以模拟
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := GetIns()
		i.id--
		fmt.Println("i.id:", i.id)
	}()
	go func() {
		defer wg.Done()
		j := GetIns()
		j.id++
		fmt.Println("j.id:", j.id)
	}()
	wg.Wait()
}
