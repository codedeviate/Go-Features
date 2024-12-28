package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4}
	defer fmt.Println("Cleaning up memory:", data)

	fmt.Println("Processing data:", data)
}
