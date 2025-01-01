package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Deferred call 1")
	}()
	defer func() {
		fmt.Println("Deferred call 2")
	}()

	fmt.Println("Starting...")
	panic("Triggering a panic")

	fmt.Println("This line never executes")
}
