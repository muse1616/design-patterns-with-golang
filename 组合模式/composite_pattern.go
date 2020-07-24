package main

import "reflect"

//员工
type Employee struct {
	name         string
	subordinates []*Employee
}

func NewEmployee(name string) *Employee {
	return &Employee{
		name:         name,
		subordinates: make([]*Employee, 0),
	}
}

// 添加 删除 得到所有子员工
func (e *Employee) add(newE *Employee) {
	e.subordinates = append(e.subordinates, newE)
}
func (e *Employee) remove(target *Employee) {
	for k, v := range e.subordinates {
		if reflect.DeepEqual(v, target) {
			e.subordinates = append(e.subordinates[:k], e.subordinates[k+1:]...)
			return
		}
	}
}
func (e *Employee) getAllSubordinates() []*Employee {
	return e.subordinates
}
