package main

import "fmt"

type State interface {
	doAction(c *Context)
}
type StartState struct {
}

func (s StartState) doAction(c *Context) {
	fmt.Println(c.id, "is in start state")
	// 设置状态
	c.setState(s)
}

type StopState struct {
}

func (s StopState) doAction(c *Context) {
	fmt.Println(c.id, "is in stop state")
	c.setState(s)
}

type Context struct {
	state State
	id    string
}

func NewContext(id string) *Context {
	return &Context{
		state: nil,
		id:    id,
	}
}
func (c *Context) setState(state State) {
	c.state = state
}
func (c *Context) getState() {
	fmt.Printf("%T \n",c.state)
}