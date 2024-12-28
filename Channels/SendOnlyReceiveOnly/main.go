package main

import "fmt"

func sendOnly(ch chan<- int) {
	ch <- 42 // Send value
}

func receiveOnly(ch <-chan int) {
	fmt.Println(<-ch) // Receive value
}

func main() {
	ch := make(chan int)

	go sendOnly(ch)
	receiveOnly(ch)
}
