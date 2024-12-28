package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function exits

	file.WriteString("Hello, Go!")
	fmt.Println("File written successfully")
}
