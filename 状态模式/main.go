package main

func main() {
	c := NewContext("1")
	start := new(StartState)
	stop := new(StopState)

	//1 is in start state
	//main.StartState
	//1 is in stop state
	//main.StopState
	start.doAction(c)
	c.getState()

	stop.doAction(c)
	c.getState()

}
