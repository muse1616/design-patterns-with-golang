package main

func main() {
	r := new(Receiver)
	//命令
	c1 := NewCommandOne(r)
	c2 := NewCommandTwo(r)
	i := NewInvoker(c1)
	//f1
	i.Do()
	i.setCommand(c2)
	//f2
	i.Do()
}
