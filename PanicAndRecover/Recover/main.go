package main

import "fmt"

func main() {
	fmt.Println("Main started")
	safeFunction()
	fmt.Println("Main ended gracefully")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("safeFunction: Doing something risky...")
	panic("An unexpected issue occurred!")
	fmt.Println("This line will never be reached")
}
