package main

import "fmt"

func createString() *string {
	s := "Garbage Collection in Go"
	return &s // Heap allocation (referenced outside the function)
}

func main() {
	str := createString()
	fmt.Println(*str) // Still in use, so not garbage collected

	str = nil // Remove reference
	// The GC will reclaim the memory at a later point
}
