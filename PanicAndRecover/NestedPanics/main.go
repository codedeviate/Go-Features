package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
			// Attempt a second panic in the same defer
			// panic("Second panic") // Uncommenting this line leads to fatal error
		}
	}()

	panic("First panic")
}
