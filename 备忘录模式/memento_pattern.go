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

//保存状态到备忘录
func (o *Originator) saveStateToMemento() *Memento {
	return NewMemento(o.state)
}
func (o *Originator) getStateFromMemento(memento *Memento) {
	o.state = memento.getState()
}

type CareTaker struct {
	mementoList []*Memento
}

func NewCareTaker() *CareTaker {
	return &CareTaker{mementoList: make([]*Memento, 0)}
}
func (c *CareTaker) add(memento *Memento) {
	c.mementoList = append(c.mementoList, memento)
}
func (c *CareTaker) get(index int) *Memento {
	return c.mementoList[index]
}
