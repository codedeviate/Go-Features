package main

import "fmt"

func main() {
	defer fmt.Println("First deferred call")
	defer fmt.Println("Second deferred call")
	defer fmt.Println("Third deferred call")
	fmt.Println("Main function executed")
}
