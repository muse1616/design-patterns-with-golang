package main

type Iterator interface {
	hasNext(a interface{}) bool
	next(a interface{}) interface{}
}

type Container interface {
	getIterator() Iterator
}

//Container实现类
type NameRepository struct {
	//姓名数字
	names    []string
	iterator NameIterator
}

func (nr *NameRepository) addNames(name []string) {
	nr.names = append(nr.names, name...)
}

// 姓名迭代器 实现Iterator接口
type NameIterator struct {
	index int
}

func (ni *NameIterator) hasNext(names interface{}) bool {
	if ni.index < len(names.([]string))-1 {
		return true
	}
	return false
}
func (ni *NameIterator) next(names interface{}) interface{} {
	ni.index++
	return (names.([]string))[ni.index]
}

func (nr *NameRepository) getIterator() Iterator {
	return &NameIterator{index: -1}
}
