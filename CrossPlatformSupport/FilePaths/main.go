package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "example"
	filename := "file.txt"

	// Cross-platform file path
	path := filepath.Join(dir, filename)
	fmt.Println("File Path:", path)

	// Creating the directory
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Creating a file
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully:", path)
}
