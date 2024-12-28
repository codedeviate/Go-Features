package main

import "fmt"

func main() {
	defer fmt.Println("This runs last")
	fmt.Println("This runs first")
}
