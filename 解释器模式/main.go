package main

import "fmt"

func main() {
	var (
		robert  = NewTerminalExpression("Robert")
		john    = NewTerminalExpression("John")
		julie   = NewTerminalExpression("Julie")
		married = NewTerminalExpression("Married")
		ex1     Expression
		ex2     Expression
	)

	ex1 = NewOrExpression(robert, john)
	ex2 = NewAndExpression(julie, married)
	// robert或者john任意一个包含即可
	fmt.Println(ex1.interpret("John"))
	// 全都需要包含
	fmt.Println(ex2.interpret("Married Julie"))
}
