package main

import "fmt"

func main() {
	ch := make(chan int, 3) // Create a buffered channel

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch) // Output: 1
	fmt.Println(<-ch) // Output: 2
	fmt.Println(<-ch) // Output: 3
}
