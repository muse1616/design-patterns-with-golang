package main

import "fmt"

func main() {
	nr:=NameRepository{}
	n:=[]string{"a","b","c","d"}
	nr.addNames(n)

	for v:=nr.getIterator();v.hasNext(nr.names);{
		fmt.Println(v.next(nr.names).(string))
	}
}