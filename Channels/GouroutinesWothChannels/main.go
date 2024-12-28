package main

import (
	"fmt"
)

func produce(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // Close the channel
}

func main() {
	ch := make(chan int)

	go produce(ch)

	for val := range ch { // Receive values until the channel is closed
		fmt.Println(val)
	}
}
