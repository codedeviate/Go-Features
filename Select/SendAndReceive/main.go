package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case ch <- 42: // Try sending
		fmt.Println("Sent 42 to channel")
	case val := <-ch: // Try receiving
		fmt.Println("Received:", val)
	default:
		fmt.Println("No operation possible")
	}
}
