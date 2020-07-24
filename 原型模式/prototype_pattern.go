package main



type Object struct {
	id int
}

func (o *Object) clone() *Object {
	// 浅拷贝
	res := *o
	return &res
}
