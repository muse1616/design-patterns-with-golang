package main

import "fmt"

// receiver
type Receiver struct {
}

// Receiver 有两个命令
func (r Receiver) f1() {
	fmt.Println("f1")
}
func (r Receiver) f2() {
	fmt.Println("f2")
}

// 命令接口
type Command interface {
	execute()
}

//实现接口
type CommandOne struct {
	r *Receiver
}
type CommandTwo struct {
	r *Receiver
}

// Constructor
func NewCommandOne(r *Receiver) *CommandOne {
	return &CommandOne{r: r}
}
func NewCommandTwo(r *Receiver) *CommandTwo {
	return &CommandTwo{r: r}
}
func (c CommandOne) execute() {
	c.r.f1()
}
func (c CommandTwo) execute() {
	c.r.f2()
}

//接受这
type Invoker struct {
	cmd Command
}

func NewInvoker(cmd Command) *Invoker {
	return &Invoker{cmd: cmd}
}
func (i *Invoker) setCommand(c Command) {
	i.cmd = c
}
func (i *Invoker)Do()  {
	i.cmd.execute()
}
