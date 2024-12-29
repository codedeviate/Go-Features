package main

import "fmt"

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func main() {
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node1.Next = node2
	node2.Previous = node1

	// Remove references
	node1 = nil
	node2 = nil

	// The GC will reclaim both nodes since they are unreachable
	fmt.Println("Circular reference handled!")
}
