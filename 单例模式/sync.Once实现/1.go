package main

import "sync"

type Singleton struct {
}

var ins *Singleton
var once sync.Once

//sync.Once内部本质上也是双重检查的方式
func GetIns() *Singleton {
	once.Do(func() {
		ins = &Singleton{}
	})
	return ins
}
