package main

import "fmt"

func main() {
	x := 5
	defer fmt.Println("Deferred value of x:", x)
	x = 10
	fmt.Println("Current value of x:", x)
}
