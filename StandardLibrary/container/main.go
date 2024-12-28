package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(10)
	l.PushBack(20)
	l.PushFront(5)

	// Iterate through the list
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
