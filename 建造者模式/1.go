package main

// 创建复杂类
type Object struct {
	// 模拟这是一个很复杂的类
	id   string
	name string
}

// getter setter
func (o *Object) setName(name string) {
	o.name = name
}
func (o *Object) setId(id string) {
	o.id = id
}

//建造者接口
type Builder interface {
	setId(id string) Builder
	setName(name string) Builder
	build() *Object
}

// 创建者类
type ObjectBuilder struct {
	object *Object
}

func (objectBuilder *ObjectBuilder) setName(name string) Builder {
	if objectBuilder.object == nil {
		objectBuilder.object = new(Object)
	}
	objectBuilder.object.setName(name)
	return objectBuilder
}

func (objectBuilder *ObjectBuilder) setId(id string) Builder {
	if objectBuilder.object == nil {
		objectBuilder.object = new(Object)
	}
	objectBuilder.object.setId(id)
	return objectBuilder
}
func (objectBuilder *ObjectBuilder) build() *Object {
	return objectBuilder.object
}

//Director实现
//Director的作用主要是规范对象的组件的创建流程
type Director struct {
	builder Builder
}

func (d Director) Create(id, name string) *Object {
	return d.builder.setId(id).setName(name).build()
}
