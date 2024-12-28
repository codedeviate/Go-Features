package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current Time:", now)

	// Formatting time
	fmt.Println("Formatted:", now.Format("2006-01-02 15:04:05"))

	// Adding time
	later := now.Add(2 * time.Hour)
	fmt.Println("Two Hours Later:", later)
}
