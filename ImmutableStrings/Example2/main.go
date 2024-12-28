package main

import (
	"fmt"
)

func main() {
	str := "Hello, World!"
	// Convert string to a slice of bytes
	bytes := []byte(str)

	// Modify the byte slice
	bytes[0] = 'h'

	// Convert back to string
	newStr := string(bytes)
	fmt.Println("Original String:", str)
	fmt.Println("Modified String:", newStr)
}
