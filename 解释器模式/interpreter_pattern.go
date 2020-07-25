package main

import "strings"

//表达式接口
type Expression interface {
	interpret(context string) bool
}
type TerminalExpression struct {
	data string
}

// Constructor
func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{data: data}
}
func (t *TerminalExpression) interpret(context string) bool {
	if strings.Contains(context, t.data) {
		return true
	}
	return false
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewOrExpression(expr1, expr2 Expression) *OrExpression {
	return &OrExpression{
		expr1: expr1,
		expr2: expr2,
	}
}
func (o *OrExpression) interpret(context string) bool {
	return o.expr1.interpret(context) || o.expr2.interpret(context)
}

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewAndExpression(expr1, expr2 Expression) *AndExpression {
	return &AndExpression{
		expr1: expr1,
		expr2: expr2,
	}
}
func (o *AndExpression) interpret(context string) bool {
	return o.expr1.interpret(context) && o.expr2.interpret(context)
}
