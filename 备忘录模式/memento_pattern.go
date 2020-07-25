package main

type Memento struct {
	// 状态
	state string
}

func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

//getter
func (m Memento) getState() string {
	return m.state
}

type Originator struct {
	state string
}

//setter and getter
func (o *Originator) getState() string {
	return o.state
}
func (o *Originator) setState(state string) {
	o.state = state
}

//生成当前状态的备忘录
func (o *Originator) generateTheMemento() *Memento {
	return NewMemento(o.state)
}
// 设置备忘录
func (o *Originator) loadStateFromMemento(memento *Memento) {
	o.state = memento.getState()
}

// 保存备忘录的结构体
type CareTaker struct {
	mementoList []*Memento
}

// Constructor
func NewCareTaker() *CareTaker {
	return &CareTaker{mementoList: make([]*Memento, 0)}
}

// 添加一条记录
func (c *CareTaker) add(memento *Memento) {
	c.mementoList = append(c.mementoList, memento)
}

// 读取记录
func (c *CareTaker) get(index int) *Memento {
	return c.mementoList[index]
}
