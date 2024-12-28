package main

import (
	"fmt"
	"os"
)

func main() {
	// Write to a file
	data := []byte("Hello, Go standard library!")
	err := os.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully")

	// Read from a file
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(content))
}
