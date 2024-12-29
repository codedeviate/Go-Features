package main

import "fmt"

func main() {
	num := new(int) // Allocated on the heap
	*num = 42
	fmt.Println("Number:", *num)

	// Remove reference
	num = nil
	// GC will reclaim the memory used by the integer at a later point
}
