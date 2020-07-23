package main

import "sync"

type Singleton struct {
	id int
}

var ins *Singleton
var mu sync.Mutex

func GetIns() *Singleton {
	// 保证线程安全
	mu.Lock()
	defer mu.Unlock()
	if ins == nil {
		ins = &Singleton{}
	}
	return ins
}
