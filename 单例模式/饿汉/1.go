package main

// 这种方式比较常用
//缺点：如果singleton创建初始化比较复杂耗时时，加载时间会延长。

type Singleton struct {
	id int
}

var ins = &Singleton{}

func GetIns() *Singleton {
	return ins
}
