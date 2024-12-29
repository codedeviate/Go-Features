package main

import (
	"fmt"
)

func readFile(fileName string) error {
	return fmt.Errorf("failed to read file %s: %w", fileName, fmt.Errorf("file not found"))
}

func main() {
	err := readFile("data.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
