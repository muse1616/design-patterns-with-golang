package main

import "fmt"

func main() {
	// 		 1
	// 		/ \
	// 	   3  4
	// 	     /
	//	    2
	e1 := NewEmployee("1")
	e2 := NewEmployee("2")
	e3 := NewEmployee("3")
	e4 := NewEmployee("4")
	e1.add(e3)
	e1.add(e4)
	e4.add(e2)

	//层次
	traverse(e1)

	// 2 3
	//for _,v:=range e1.getAllSubordinates(){
	//	fmt.Println(v.name)
	//}

}
func traverse(e1 *Employee) {
	if e1 == nil {
		return
	}
	queue := make([]*Employee, 0)
	queue = append(queue, e1)
	for len(queue) != 0 {
		//	取第一个元素
		node := queue[0]

		fmt.Print(node.name, "\t")
		// 子节点加入队列
		for _, v := range node.subordinates {
			queue = append(queue, v)
		}
		//	出列
		queue = queue[1:]
	}

}
