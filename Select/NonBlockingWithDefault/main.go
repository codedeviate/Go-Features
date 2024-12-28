package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	default:
		fmt.Println("No data received")
	}
}
