package main

import "fmt"

func main() {
	originator:=Originator{}
	careTaker:=NewCareTaker()
	originator.setState("state #1")
	originator.setState("state #2")
	// 保存2
	careTaker.add(originator.saveStateToMemento())
	originator.setState("state #3")
	//保存3
	careTaker.add(originator.saveStateToMemento())
	originator.setState("state #4")

	fmt.Printf("当前状态:%v\n",originator.getState())
	originator.getStateFromMemento(careTaker.mementoList[0])
	fmt.Println("第一次保存的状态:",originator.getState())
	originator.getStateFromMemento(careTaker.mementoList[1])
	fmt.Println("第二次保存的状态:",originator.getState())
}