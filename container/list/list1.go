package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e1 := l.PushBack(4)
	e2 := l.PushFront(1)
	l.InsertBefore(3, e1)
	l.InsertAfter(2, e2)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}