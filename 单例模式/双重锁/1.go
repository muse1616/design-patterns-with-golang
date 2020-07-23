package main

import "sync"

type Singleton struct {
	id int
}

var ins *Singleton
var mu sync.Mutex

// 这样避免每次都加锁 而是只有未初始化之前才会加锁判断
func GetIns() *Singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			ins = &Singleton{}
		}
	}
	return ins
}
